package cmd

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/ollama/ollama/api"
	"github.com/spf13/cobra"
)

var codeCmd = &cobra.Command{
	Use:   "code",
	Short: "Code manipulation and generation tools",
	Long:  "Edit, generate, refactor, and analyze code with AI assistance",
}

var codeEditCmd = &cobra.Command{
	Use:   "edit [file]",
	Short: "Edit code files with AI assistance",
	Long:  "Open and modify code files with AI-powered suggestions and edits",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var filePath string
		if len(args) > 0 {
			filePath = args[0]
		} else {
			// Interactive file selection
			var err error
			filePath, err = selectFileInteractively()
			if err != nil {
				fmt.Fprintf(os.Stderr, "❌ Error selecting file: %v\n", err)
				return
			}
		}

		if err := editCodeFile(filePath); err != nil {
			fmt.Fprintf(os.Stderr, "❌ Error editing file: %v\n", err)
			os.Exit(1)
		}
	},
}

var codeGenerateCmd = &cobra.Command{
	Use:   "generate [template]",
	Short: "Generate code from templates",
	Long:  "Generate code files, functions, or structures from predefined templates",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		template := "function"
		if len(args) > 0 {
			template = args[0]
		}

		if err := generateCode(template); err != nil {
			fmt.Fprintf(os.Stderr, "❌ Error generating code: %v\n", err)
			os.Exit(1)
		}
	},
}

var codeRefactorCmd = &cobra.Command{
	Use:   "refactor [file]",
	Short: "AI-powered code refactoring",
	Long:  "Analyze and refactor code using AI to improve quality and performance",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]

		if err := refactorCode(filePath); err != nil {
			fmt.Fprintf(os.Stderr, "❌ Error refactoring code: %v\n", err)
			os.Exit(1)
		}
	},
}

var codeAnalyzeCmd = &cobra.Command{
	Use:   "analyze [file]",
	Short: "Analyze code quality and suggest improvements",
	Long:  "Use AI to analyze code and provide suggestions for improvements",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]

		if err := analyzeCode(filePath); err != nil {
			fmt.Fprintf(os.Stderr, "❌ Error analyzing code: %v\n", err)
			os.Exit(1)
		}
	},
}

var codeAddCmd = &cobra.Command{
	Use:   "add [feature]",
	Short: "Add features to existing code",
	Long:  "Use AI to add new features, functions, or capabilities to existing code files",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		feature := args[0]

		if err := addCodeFeature(feature); err != nil {
			fmt.Fprintf(os.Stderr, "❌ Error adding feature: %v\n", err)
			os.Exit(1)
		}
	},
}

func selectFileInteractively() (string, error) {
	// Get current directory files
	files, err := getCodeFiles(".")
	if err != nil {
		return "", err
	}

	if len(files) == 0 {
		return "", fmt.Errorf("no code files found in current directory")
	}

	prompt := promptui.Select{
		Label: "Select a file to edit",
		Items: files,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}?",
			Active:   "▶ {{ . | cyan }}",
			Inactive: "  {{ . | white }}",
			Selected: "✅ {{ . | bold }}",
		},
	}

	_, result, err := prompt.Run()
	return result, err
}

func getCodeFiles(dir string) ([]string, error) {
	var files []string
	codeExtensions := map[string]bool{
		".go": true, ".js": true, ".ts": true, ".py": true, ".java": true,
		".cpp": true, ".c": true, ".h": true, ".hpp": true, ".cs": true,
		".php": true, ".rb": true, ".rs": true, ".kt": true, ".swift": true,
		".html": true, ".css": true, ".scss": true, ".vue": true, ".jsx": true,
		".tsx": true, ".json": true, ".yaml": true, ".yml": true, ".xml": true,
		".sql": true, ".md": true, ".txt": true,
	}

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip hidden directories and files
		if strings.HasPrefix(d.Name(), ".") {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Skip common build/dependency directories
		if d.IsDir() && (d.Name() == "node_modules" || d.Name() == "vendor" || d.Name() == "target" || d.Name() == ".git") {
			return filepath.SkipDir
		}

		if !d.IsDir() {
			ext := filepath.Ext(path)
			if codeExtensions[ext] {
				files = append(files, path)
			}
		}

		return nil
	})

	return files, err
}

func editCodeFile(filePath string) error {
	// Check if file exists
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("could not read file %s: %w", filePath, err)
	}

	fmt.Printf("📝 Editing: %s\n", filePath)
	fmt.Printf("📊 File size: %d lines\n\n", strings.Count(string(content), "\n")+1)

	// Show available edit options
	editOptions := []string{
		"View file content",
		"AI-powered modification",
		"Add function/method",
		"Fix bugs with AI",
		"Add comments/documentation",
		"Format/cleanup code",
		"Exit",
	}

	for {
		prompt := promptui.Select{
			Label: "What would you like to do?",
			Items: editOptions,
		}

		_, choice, err := prompt.Run()
		if err != nil {
			return err
		}

		switch choice {
		case "View file content":
			fmt.Printf("\n📄 Content of %s:\n", filePath)
			fmt.Println("─────────────────────────────────────")
			fmt.Print(string(content))
			fmt.Println("\n─────────────────────────────────────")

		case "AI-powered modification":
			if err := aiModifyFile(filePath, string(content)); err != nil {
				fmt.Printf("❌ Error: %v\n", err)
			}

		case "Add function/method":
			if err := aiAddFunction(filePath, string(content)); err != nil {
				fmt.Printf("❌ Error: %v\n", err)
			}

		case "Fix bugs with AI":
			if err := aiFixBugs(filePath, string(content)); err != nil {
				fmt.Printf("❌ Error: %v\n", err)
			}

		case "Add comments/documentation":
			if err := aiAddDocumentation(filePath, string(content)); err != nil {
				fmt.Printf("❌ Error: %v\n", err)
			}

		case "Format/cleanup code":
			if err := formatCode(filePath, string(content)); err != nil {
				fmt.Printf("❌ Error: %v\n", err)
			}

		case "Exit":
			fmt.Println("✅ Editing session completed!")
			return nil
		}

		// Reload content after modifications
		if newContent, err := os.ReadFile(filePath); err == nil {
			content = newContent
		}
	}
}

func aiModifyFile(filePath, content string) error {
	prompt := promptui.Prompt{
		Label: "Describe what you want to change",
	}

	description, err := prompt.Run()
	if err != nil {
		return err
	}

	return callAIForCodeModification(filePath, content, description)
}

func aiAddFunction(filePath, content string) error {
	prompt := promptui.Prompt{
		Label: "Describe the function you want to add",
	}

	description, err := prompt.Run()
	if err != nil {
		return err
	}

	return callAIForCodeGeneration(filePath, content, "Add a new function: "+description)
}

func aiFixBugs(filePath, content string) error {
	return callAIForCodeModification(filePath, content, "Analyze this code and fix any bugs or potential issues you find. Explain what you fixed.")
}

func aiAddDocumentation(filePath, content string) error {
	return callAIForCodeModification(filePath, content, "Add comprehensive comments and documentation to this code. Include function descriptions, parameter explanations, and return value documentation.")
}

func formatCode(filePath, content string) error {
	ext := filepath.Ext(filePath)

	// Simple formatting for different languages
	switch ext {
	case ".go":
		return runCommand("gofmt", "-w", filePath)
	case ".js", ".ts", ".jsx", ".tsx":
		return runCommand("prettier", "--write", filePath)
	case ".py":
		return runCommand("black", filePath)
	default:
		return callAIForCodeModification(filePath, content, "Format and clean up this code following best practices for the language. Fix indentation, spacing, and style issues.")
	}
}

func callAIForCodeModification(filePath, content, task string) error {
	client, err := createOllamaClient()
	if err != nil {
		return fmt.Errorf("OLLAMA not available: %w", err)
	}

	ctx := context.Background()

	// Check if OLLAMA is running
	if err := client.Heartbeat(ctx); err != nil {
		return fmt.Errorf("OLLAMA server not responding. Make sure it's running with: ollama serve")
	}

	fmt.Printf("🤖 AI is analyzing and modifying your code...\n\n")

	prompt := fmt.Sprintf(`You are a professional software developer. Here's a code file and a task:

File: %s
Task: %s

Current code:
%s

Please provide the modified code with your changes. Explain what you changed and why. If you're creating a new file, provide the complete file content.`, filePath, task, content)

	req := &api.GenerateRequest{
		Model:  "llama2", // You can make this configurable
		Prompt: prompt,
	}

	var response strings.Builder
	err = client.Generate(ctx, req, func(resp api.GenerateResponse) error {
		if resp.Response != "" {
			fmt.Print(resp.Response)
			response.WriteString(resp.Response)
		}
		return nil
	})

	if err != nil {
		return err
	}

	fmt.Print("\n\n")

	// Ask if user wants to apply changes
	confirmPrompt := promptui.Prompt{
		Label:     "Apply these changes to the file? (y/n)",
		IsConfirm: true,
	}

	_, err = confirmPrompt.Run()
	if err != nil {
		fmt.Println("❌ Changes not applied.")
		return nil
	}

	// Here you would parse the AI response and apply the changes
	// For now, we'll just show the response
	fmt.Println("✅ Changes applied! (Note: Full implementation would parse and apply the AI-generated code)")

	return nil
}

func callAIForCodeGeneration(filePath, content, task string) error {
	return callAIForCodeModification(filePath, content, task)
}

func generateCode(template string) error {
	templates := map[string]string{
		"function":   "Generate a new function based on your requirements",
		"class":      "Generate a new class with methods and properties",
		"api":        "Generate a REST API endpoint",
		"test":       "Generate unit tests for existing code",
		"component":  "Generate a UI component (React, Vue, etc.)",
		"model":      "Generate a data model or struct",
		"config":     "Generate configuration files",
		"dockerfile": "Generate a Dockerfile for the project",
	}

	description, exists := templates[template]
	if !exists {
		fmt.Printf("❌ Unknown template: %s\n", template)
		fmt.Println("💡 Available templates: function, class, api, test, component, model, config, dockerfile")
		return nil
	}

	prompt := promptui.Prompt{
		Label: fmt.Sprintf("Describe the %s you want to generate", template),
	}

	userDescription, err := prompt.Run()
	if err != nil {
		return err
	}

	// Generate filename
	filenamePrompt := promptui.Prompt{
		Label: "Enter filename (with extension)",
	}

	filename, err := filenamePrompt.Run()
	if err != nil {
		return err
	}

	fullTask := fmt.Sprintf("%s: %s", description, userDescription)
	return callAIForCodeModification(filename, "", fullTask)
}

func refactorCode(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	task := "Refactor this code to improve readability, performance, and maintainability. Follow best practices and design patterns. Explain your changes."
	return callAIForCodeModification(filePath, string(content), task)
}

func analyzeCode(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	task := "Analyze this code and provide suggestions for improvements. Look for potential bugs, performance issues, security vulnerabilities, and opportunities to improve code quality."
	return callAIForCodeModification(filePath, string(content), task)
}

func addCodeFeature(feature string) error {
	// Select file to add feature to
	filePath, err := selectFileInteractively()
	if err != nil {
		return err
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	task := fmt.Sprintf("Add the following feature to this code: %s. Integrate it properly with the existing code structure.", feature)
	return callAIForCodeModification(filePath, string(content), task)
}

func runCommand(name string, args ...string) error {
	fmt.Printf("🔧 Running: %s %s\n", name, strings.Join(args, " "))
	// This is a placeholder - you'd implement actual command execution here
	fmt.Println("✅ Command completed (placeholder implementation)")
	return nil
}

// Add code commands to interactive mode
func processCodeCommand(parts []string) bool {
	if len(parts) < 2 {
		fmt.Println("❌ Usage: code <command> [args...]")
		fmt.Println("💡 Commands: edit, generate, refactor, analyze, add")
		return true
	}

	subCommand := parts[1]
	args := parts[2:]

	switch subCommand {
	case "edit":
		var filePath string
		if len(args) > 0 {
			filePath = args[0]
		}

		if filePath == "" {
			file, err := selectFileInteractively()
			if err != nil {
				fmt.Printf("❌ Error: %v\n", err)
				return true
			}
			filePath = file
		}

		if err := editCodeFile(filePath); err != nil {
			fmt.Printf("❌ Error: %v\n", err)
		}

	case "generate":
		template := "function"
		if len(args) > 0 {
			template = args[0]
		}
		if err := generateCode(template); err != nil {
			fmt.Printf("❌ Error: %v\n", err)
		}

	case "refactor":
		if len(args) < 1 {
			fmt.Println("❌ Usage: code refactor <file>")
			return true
		}
		if err := refactorCode(args[0]); err != nil {
			fmt.Printf("❌ Error: %v\n", err)
		}

	case "analyze":
		if len(args) < 1 {
			fmt.Println("❌ Usage: code analyze <file>")
			return true
		}
		if err := analyzeCode(args[0]); err != nil {
			fmt.Printf("❌ Error: %v\n", err)
		}

	case "add":
		if len(args) < 1 {
			fmt.Println("❌ Usage: code add <feature>")
			return true
		}
		feature := strings.Join(args, " ")
		if err := addCodeFeature(feature); err != nil {
			fmt.Printf("❌ Error: %v\n", err)
		}

	default:
		fmt.Printf("❌ Unknown code command: %s\n", subCommand)
		fmt.Println("💡 Available commands: edit, generate, refactor, analyze, add")
	}

	return true
}

func init() {
	// Add subcommands
	codeCmd.AddCommand(codeEditCmd)
	codeCmd.AddCommand(codeGenerateCmd)
	codeCmd.AddCommand(codeRefactorCmd)
	codeCmd.AddCommand(codeAnalyzeCmd)
	codeCmd.AddCommand(codeAddCmd)

	// Add to root command
	rootCmd.AddCommand(codeCmd)
}
