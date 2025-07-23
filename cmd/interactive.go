package cmd

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var interactiveCmd = &cobra.Command{
	Use:     "interactive",
	Short:   "Start interactive mode with command input box",
	Long:    "Enter interactive mode where you can type commands in a text box interface",
	Aliases: []string{"i", "shell"},
	Run: func(cmd *cobra.Command, args []string) {
		startInteractiveMode()
	},
}

func startInteractiveMode() {
	fmt.Printf("\n🎯 Welcome to DevLama Interactive Mode!\n")
	fmt.Printf("Type 'help' for available commands or 'exit' to quit.\n\n")

	for {
		prompt := promptui.Prompt{
			Label: "DevLama",
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
				fmt.Println("\n👋 Goodbye!")
				return
			}
			fmt.Printf("❌ Error: %v\n", err)
			continue
		}

		// Process the command
		if !processInteractiveCommand(strings.TrimSpace(input)) {
			break
		}
	}
}

func processInteractiveCommand(input string) bool {
	if input == "" {
		return true
	}

	parts := strings.Fields(input)
	command := parts[0]

	switch command {
	case "exit", "quit", "q":
		fmt.Println("👋 Goodbye!")
		return false

	case "help", "h":
		showInteractiveHelp()

	case "clear", "cls":
		clearScreen()

	case "hello":
		if len(parts) > 1 {
			name := strings.Join(parts[1:], " ")
			fmt.Printf("👋 Hello, %s!\n", name)
		} else {
			fmt.Println("👋 Hello, World!")
		}

	case "version", "v":
		fmt.Printf("📦 DevLama v%s\n", version)

	case "time", "date":
		showCurrentTime()

	case "banner":
		showBanner()

	case "open":
		if len(parts) > 1 {
			target := parts[1]
			if err := openTarget(target); err != nil {
				fmt.Printf("❌ Error opening %s: %v\n", target, err)
			} else {
				fmt.Printf("✅ Opened: %s\n", target)
			}
		} else {
			fmt.Println("❌ Usage: open <file|directory|url>")
		}

	case "echo":
		if len(parts) > 1 {
			message := strings.Join(parts[1:], " ")
			fmt.Printf("🔊 %s\n", message)
		} else {
			fmt.Println("❌ Usage: echo <message>")
		}

	case "ollama":
		return processOllamaCommand(parts)

	case "code":
		return processCodeCommand(parts)

	default:
		fmt.Printf("❌ Unknown command: %s\n", command)
		fmt.Println("💡 Type 'help' for available commands")
	}

	return true
}

func showInteractiveHelp() {
	fmt.Println("\n📋 Available Commands:")
	fmt.Println("  • hello [name]        - Greet someone")
	fmt.Println("  • version (v)         - Show version")
	fmt.Println("  • open <target>       - Open file/directory/URL")
	fmt.Println("  • echo <message>      - Echo a message")
	fmt.Println("  • time                - Show current time")
	fmt.Println("  • banner              - Show DevLama banner")
	fmt.Println("  • clear (cls)         - Clear screen")
	fmt.Println()
	fmt.Println("🤖 AI Commands:")
	fmt.Println("  • ollama chat [model] - Chat with AI (e.g., llama2)")
	fmt.Println("  • ollama list         - List available AI models")
	fmt.Println("  • ollama pull <model> - Download AI model")
	fmt.Println()
	fmt.Println("💻 Code Commands:")
	fmt.Println("  • code edit [file]    - Edit code with AI assistance")
	fmt.Println("  • code generate       - Generate code templates")
	fmt.Println("  • code refactor <file>- AI-powered refactoring")
	fmt.Println("  • code analyze <file> - Analyze code quality")
	fmt.Println("  • code add <feature>  - Add features to code")
	fmt.Println()
	fmt.Println("  • help (h)            - Show this help")
	fmt.Println("  • exit (quit, q)      - Exit interactive mode")
	fmt.Println()
}

func clearScreen() {
	fmt.Print("\033[2J\033[H")
}

func showCurrentTime() {
	fmt.Printf("🕒 Current time: %s\n", getCurrentTime())
}

func init() {
	rootCmd.AddCommand(interactiveCmd)
}
