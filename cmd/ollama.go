package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/ollama/ollama/api"
	"github.com/spf13/cobra"
)

var ollamaCmd = &cobra.Command{
	Use:   "ollama",
	Short: "Interact with OLLAMA AI models",
	Long:  "Chat with AI models, list available models, and manage OLLAMA locally",
}

var ollamaChatCmd = &cobra.Command{
	Use:   "chat [model]",
	Short: "Start an interactive chat with an OLLAMA model",
	Long:  "Start an interactive chat session with the specified OLLAMA model",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		modelName := "llama2" // default model
		if len(args) > 0 {
			modelName = args[0]
		}

		if err := startOllamaChat(modelName); err != nil {
			fmt.Fprintf(os.Stderr, "❌ Error starting chat: %v\n", err)
			os.Exit(1)
		}
	},
}

var ollamaListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available OLLAMA models",
	Long:  "Display all locally available OLLAMA models",
	Run: func(cmd *cobra.Command, args []string) {
		if err := listOllamaModels(); err != nil {
			fmt.Fprintf(os.Stderr, "❌ Error listing models: %v\n", err)
			os.Exit(1)
		}
	},
}

var ollamaGenerateCmd = &cobra.Command{
	Use:   "generate [model] [prompt]",
	Short: "Generate text with an OLLAMA model",
	Long:  "Generate a one-time response from an OLLAMA model",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		modelName := args[0]
		prompt := strings.Join(args[1:], " ")

		if err := generateOllamaText(modelName, prompt); err != nil {
			fmt.Fprintf(os.Stderr, "❌ Error generating text: %v\n", err)
			os.Exit(1)
		}
	},
}

var ollamaPullCmd = &cobra.Command{
	Use:   "pull [model]",
	Short: "Download an OLLAMA model",
	Long:  "Download and install a model from the OLLAMA registry",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		modelName := args[0]

		if err := pullOllamaModel(modelName); err != nil {
			fmt.Fprintf(os.Stderr, "❌ Error pulling model: %v\n", err)
			os.Exit(1)
		}
	},
}

func createOllamaClient() (*api.Client, error) {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		return nil, fmt.Errorf("failed to create OLLAMA client: %w (is OLLAMA running?)", err)
	}
	return client, nil
}

func startOllamaChat(modelName string) error {
	client, err := createOllamaClient()
	if err != nil {
		return err
	}

	// Check if OLLAMA is running
	ctx := context.Background()
	if err := client.Heartbeat(ctx); err != nil {
		return fmt.Errorf("OLLAMA server not responding: %w\nMake sure OLLAMA is running with: ollama serve", err)
	}

	fmt.Printf("🤖 Starting chat with %s\n", modelName)
	fmt.Printf("💡 Type 'exit', 'quit', or press Ctrl+C to end the chat\n\n")

	messages := []api.Message{}

	for {
		prompt := promptui.Prompt{
			Label: "You",
			Templates: &promptui.PromptTemplates{
				Prompt:  "{{ . | cyan }}{{ \":\" | cyan }} ",
				Valid:   "{{ . | cyan }}{{ \":\" | cyan }} ",
				Invalid: "{{ . | red }}{{ \":\" | red }} ",
				Success: "{{ . | bold }}{{ \":\" | cyan }} ",
			},
		}

		input, err := prompt.Run()
		if err != nil {
			if err == promptui.ErrInterrupt {
				fmt.Println("\n👋 Chat ended!")
				return nil
			}
			return err
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		if input == "exit" || input == "quit" {
			fmt.Println("👋 Chat ended!")
			return nil
		}

		// Add user message
		messages = append(messages, api.Message{
			Role:    "user",
			Content: input,
		})

		// Generate response
		fmt.Print("🤖 ")

		req := &api.ChatRequest{
			Model:    modelName,
			Messages: messages,
		}

		var assistantResponse strings.Builder

		err = client.Chat(ctx, req, func(resp api.ChatResponse) error {
			if resp.Message.Content != "" {
				fmt.Print(resp.Message.Content)
				assistantResponse.WriteString(resp.Message.Content)
			}
			return nil
		})

		if err != nil {
			fmt.Printf("\n❌ Error: %v\n\n", err)
			continue
		}

		fmt.Print("\n\n")

		// Add assistant response to messages
		messages = append(messages, api.Message{
			Role:    "assistant",
			Content: assistantResponse.String(),
		})
	}
}

func listOllamaModels() error {
	client, err := createOllamaClient()
	if err != nil {
		return err
	}

	ctx := context.Background()
	resp, err := client.List(ctx)
	if err != nil {
		return err
	}

	if len(resp.Models) == 0 {
		fmt.Println("📦 No models found locally")
		fmt.Println("💡 Pull a model with: devlama ollama pull llama2")
		return nil
	}

	fmt.Printf("📦 Available OLLAMA Models (%d):\n\n", len(resp.Models))

	for _, model := range resp.Models {
		fmt.Printf("🤖 %s\n", model.Name)
		fmt.Printf("   Size: %s\n", formatBytes(model.Size))
		fmt.Printf("   Modified: %s\n", model.ModifiedAt.Format("2006-01-02 15:04:05"))
		if model.Details.Family != "" {
			fmt.Printf("   Family: %s\n", model.Details.Family)
		}
		fmt.Println()
	}

	return nil
}

func generateOllamaText(modelName, prompt string) error {
	client, err := createOllamaClient()
	if err != nil {
		return err
	}

	ctx := context.Background()

	fmt.Printf("🤖 Generating response with %s...\n\n", modelName)

	req := &api.GenerateRequest{
		Model:  modelName,
		Prompt: prompt,
	}

	err = client.Generate(ctx, req, func(resp api.GenerateResponse) error {
		if resp.Response != "" {
			fmt.Print(resp.Response)
		}
		return nil
	})

	if err != nil {
		return err
	}

	fmt.Print("\n\n")
	return nil
}

func pullOllamaModel(modelName string) error {
	client, err := createOllamaClient()
	if err != nil {
		return err
	}

	ctx := context.Background()

	fmt.Printf("📥 Pulling model: %s\n", modelName)

	req := &api.PullRequest{
		Model: modelName,
	}

	err = client.Pull(ctx, req, func(resp api.ProgressResponse) error {
		if resp.Total > 0 {
			progress := float64(resp.Completed) / float64(resp.Total) * 100
			fmt.Printf("\r📥 %s: %.1f%% (%s/%s)",
				resp.Status,
				progress,
				formatBytes(resp.Completed),
				formatBytes(resp.Total))
		} else {
			fmt.Printf("\r📥 %s", resp.Status)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("\n❌ Error: %v\n", err)
		return err
	}

	fmt.Printf("\n✅ Successfully pulled %s\n", modelName)
	return nil
}

func formatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// Add OLLAMA support to interactive mode
func processOllamaCommand(parts []string) bool {
	if len(parts) < 2 {
		fmt.Println("❌ Usage: ollama <command> [args...]")
		fmt.Println("💡 Commands: chat, list, generate, pull")
		return true
	}

	subCommand := parts[1]
	args := parts[2:]

	switch subCommand {
	case "chat":
		modelName := "llama2"
		if len(args) > 0 {
			modelName = args[0]
		}
		fmt.Printf("🚀 Starting OLLAMA chat with %s...\n", modelName)
		if err := startOllamaChat(modelName); err != nil {
			fmt.Printf("❌ Error: %v\n", err)
		}

	case "list":
		if err := listOllamaModels(); err != nil {
			fmt.Printf("❌ Error: %v\n", err)
		}

	case "generate":
		if len(args) < 2 {
			fmt.Println("❌ Usage: ollama generate <model> <prompt>")
			return true
		}
		modelName := args[0]
		prompt := strings.Join(args[1:], " ")
		if err := generateOllamaText(modelName, prompt); err != nil {
			fmt.Printf("❌ Error: %v\n", err)
		}

	case "pull":
		if len(args) < 1 {
			fmt.Println("❌ Usage: ollama pull <model>")
			return true
		}
		if err := pullOllamaModel(args[0]); err != nil {
			fmt.Printf("❌ Error: %v\n", err)
		}

	default:
		fmt.Printf("❌ Unknown OLLAMA command: %s\n", subCommand)
		fmt.Println("💡 Available commands: chat, list, generate, pull")
	}

	return true
}

func init() {
	// Add subcommands
	ollamaCmd.AddCommand(ollamaChatCmd)
	ollamaCmd.AddCommand(ollamaListCmd)
	ollamaCmd.AddCommand(ollamaGenerateCmd)
	ollamaCmd.AddCommand(ollamaPullCmd)

	// Add to root command
	rootCmd.AddCommand(ollamaCmd)
}
