# DevLama - AI-Powered Developer Productivity CLI

```
██████╗ ███████╗██╗   ██╗██╗      █████╗ ███╗   ███╗ █████╗ 
██╔══██╗██╔════╝██║   ██║██║     ██╔══██╗████╗ ████║██╔══██╗
██║  ██║█████╗  ██║   ██║██║     ███████║██╔████╔██║███████║
██║  ██║██╔══╝  ╚██╗ ██╔╝██║     ██╔══██║██║╚██╔╝██║██╔══██║
██████╔╝███████╗ ╚████╔╝ ███████╗██║  ██║██║ ╚═╝ ██║██║  ██║
╚═════╝ ╚══════╝  ╚═══╝  ╚══════╝╚═╝  ╚═╝╚═╝     ╚═╝╚═╝  ╚═╝
```

**The Ultimate AI-Powered Command-Line Tool for Developers**

DevLama transforms your terminal into an intelligent development environment where you can chat with AI, manipulate code, and boost productivity—all through beautiful, interactive commands.

---

## The Vision

**DevLama bridges the gap between AI and your development workflow.** Instead of switching between your editor, browser, and AI tools, DevLama brings everything into your terminal:

- **Chat with AI models** locally via OLLAMA
- **Edit and generate code** with AI assistance  
- **Beautiful interactive interface** with text input boxes
- **Instant productivity** without leaving your terminal
- **Extensible architecture** for custom workflows

Think of it as **"Claude/ChatGPT for your terminal"** but with the power to actually modify your codebase.

---

## Core Features

### AI Integration (OLLAMA-Powered)
- **Interactive AI Chat**: Real-time conversations with local LLMs
- **Code Analysis**: AI-powered code review and suggestions
- **Smart Code Generation**: Create functions, classes, APIs with natural language
- **Bug Detection**: Automatic identification and fixing of code issues
- **Documentation Generation**: AI-generated comments and docs

### 💻 **Code Manipulation**
- **AI Code Editor**: Modify files with intelligent suggestions
- **Template Generation**: Pre-built templates for common patterns
- **Refactoring Engine**: AI-driven code improvements
- **Multi-Language Support**: Go, Python, JavaScript, Java, C++, and more
- **Safe Operations**: Always confirm before making changes

### Beautiful User Experience
- **Colorful ASCII Art**: Professional branding and visual appeal
- **Interactive Text Boxes**: Modern input interfaces with auto-completion
- **Progress Indicators**: Real-time feedback for long operations
- **Organized Help System**: Comprehensive documentation built-in
- **Cross-Platform**: Works on Windows, macOS, and Linux

---

## Quick Start

### Prerequisites
- Go 1.24.5+ - [Download Go](https://golang.org/dl/)
- OLLAMA - [Download OLLAMA](https://ollama.ai) for AI features

### Installation

```bash
# Clone the repository
git clone https://github.com/nitink23/devlama.git
cd devlama

# Build DevLama
go build -o devlama.exe  # Windows
# go build -o devlama     # macOS/Linux

# Test it works
./devlama.exe --help
```

### Setup AI Features (OLLAMA)

```bash
# 1. Install OLLAMA from https://ollama.ai

# 2. Start OLLAMA service
ollama serve

# 3. Download your first AI model
ollama pull llama2

# 4. Test AI integration
./devlama.exe ollama chat llama2
```

---

## Usage Examples

### Interactive Mode (Recommended)
```bash
# Start the beautiful interactive interface
./devlama.exe interactive

# Try these commands in the text box:
> hello DevLama
> ollama chat llama2
> code edit main.go
> code generate function
> banner
> help
```

### AI Commands
```bash
# Chat with AI models
./devlama.exe ollama chat llama2
./devlama.exe ollama list
./devlama.exe ollama pull codellama

# Generate text with AI
./devlama.exe ollama generate llama2 "Write a Go function to reverse a string"
```

### 💻 **Code Manipulation**
```bash
# Edit code files with AI assistance
./devlama.exe code edit main.go

# Generate new code from templates
./devlama.exe code generate function
./devlama.exe code generate api
./devlama.exe code generate test

# Analyze and improve existing code
./devlama.exe code analyze cmd/root.go
./devlama.exe code refactor utils.go

# Add features to existing files
./devlama.exe code add "user authentication system"
```

### Utility Commands
```bash
# Open files and URLs
./devlama.exe open .
./devlama.exe open https://github.com

# Show system information
./devlama.exe version
./devlama.exe

# Interactive help
./devlama.exe help
```

---

## Architecture & Design

### Command Structure
```
devlama/
├── main.go              # Application entry point
├── cmd/                 # Command definitions
│   ├── root.go          # Root command with banner
│   ├── interactive.go   # Interactive text box interface
│   ├── ollama.go        # AI model integration
│   ├── code.go          # Code manipulation commands
│   ├── open.go          # File/URL opener
│   ├── version.go       # Version and system info
│   └── hello.go         # Sample greeting command
├── go.mod               # Dependencies
├── TESTING.md           # Comprehensive testing guide
└── README.md            # This file
```

### Technology Stack
- **Framework**: [Cobra CLI](https://github.com/spf13/cobra) for command structure
- **AI Integration**: [OLLAMA API](https://github.com/ollama/ollama) for local LLMs  
- **UI Components**: [PromptUI](https://github.com/manifoldco/promptui) for interactive interfaces
- **Language**: Go 1.24.5 for performance and cross-platform support

---

## Use Cases

### For Individual Developers
- **Code Review**: Get AI feedback on your code quality
- **Learning**: Ask AI to explain complex code patterns
- **Debugging**: Let AI identify and fix bugs
- **Documentation**: Auto-generate comments and docs
- **Prototyping**: Quickly generate boilerplate code

### For Development Teams
- **Code Standards**: Enforce consistent coding practices
- **Onboarding**: Help new developers understand codebases
- **Refactoring**: Safely improve legacy code
- **API Development**: Generate REST endpoints and models
- **Testing**: Create comprehensive test suites

### For DevOps/Infrastructure
- **Configuration**: Generate Docker, Kubernetes configs
- **Scripts**: Create automation and deployment scripts
- **Monitoring**: Generate logging and alerting code
- **Documentation**: Maintain up-to-date infrastructure docs

---

## Advanced Configuration

### Custom AI Models
```bash
# Use different models for different tasks
./devlama.exe ollama pull codellama:13b    # For coding
./devlama.exe ollama pull mistral          # For general chat
./devlama.exe ollama pull deepseek-coder   # For code analysis
```

### Environment Variables
```bash
export OLLAMA_HOST=http://localhost:11434  # Custom OLLAMA server
export DEVLAMA_DEFAULT_MODEL=codellama     # Default AI model
export DEVLAMA_VERBOSE=true                # Enable verbose logging
```

### Integration with IDEs
```bash
# VS Code: Add as terminal task
# Vim: Use as external command
# IntelliJ: Configure as external tool
```

---

## Interactive Mode Deep Dive

The **interactive mode** is DevLama's crown jewel - a beautiful text-based interface that feels modern and responsive:

### Features
- **Smart Auto-completion**: Tab completion for commands and files
- **Command History**: Navigate previous commands with arrow keys
- **Colored Output**: Syntax highlighting and emoji indicators
- **Progress Bars**: Visual feedback for long-running operations
- **Error Handling**: Friendly error messages with suggestions

### Available Commands in Interactive Mode

#### AI Commands
```
ollama chat [model]     - Start AI conversation
ollama list            - Show available models  
ollama pull <model>    - Download new models
ollama generate <prompt> - One-time text generation
```

#### Code Commands
```
code edit [file]       - Edit files with AI help
code generate [type]   - Create new code
code refactor <file>   - Improve existing code
code analyze <file>    - Get code insights
code add <feature>     - Add functionality
```

#### Utility Commands
```
hello [name]          - Friendly greeting
open <target>         - Open files/URLs
time                  - Show current time
banner                - Display DevLama logo
clear                 - Clear screen
version               - Show version info
help                  - Show all commands
exit                  - Leave interactive mode
```

---

## Testing Guide

We've created a comprehensive testing guide! See **[TESTING.md](TESTING.md)** for:
- Step-by-step testing instructions
- OLLAMA setup verification
- Feature validation checklist
- Troubleshooting common issues
- Performance benchmarks

### Quick Test
```bash
# Test basic functionality
./devlama.exe version
./devlama.exe hello --name "Tester"

# Test AI integration (requires OLLAMA)
./devlama.exe ollama list

# Test interactive mode
./devlama.exe interactive
```

---

## Contributing

We welcome contributions! Here's how to get started:

### Development Setup
```bash
# Fork and clone the repo
git clone https://github.com/yourusername/devlama.git
cd devlama

# Install dependencies
go mod tidy

# Make your changes
# Test thoroughly
go build -o devlama.exe
./devlama.exe --help

# Submit a pull request
```

### Adding New Commands
1. Create a new file in `cmd/` directory
2. Follow the existing command patterns
3. Add to interactive mode in `cmd/interactive.go`
4. Update help text and documentation
5. Add tests and examples

### Contribution Areas
- **AI Features**: New OLLAMA integrations
- **Code Tools**: Additional programming languages
- **UI/UX**: Enhanced interactive experiences  
- **Documentation**: Tutorials and guides
- **Testing**: Automated test suites
- **DevOps**: CI/CD and deployment

---

## Why DevLama?

### The Problem
Modern development involves constantly switching between:
- Code editors for writing
- Browsers for research  
- AI tools for assistance
- Terminals for commands
- Documentation for references

### The Solution
DevLama unifies these workflows into a single, intelligent command-line interface:
- **AI-First**: Get help without leaving your terminal
- **Code-Aware**: Understands your project structure
- **Interactive**: Beautiful, modern CLI experience
- **Extensible**: Easy to customize and extend
- **Local**: Your code never leaves your machine

### The Result
- **10x faster** development workflow
- **Smarter** code with AI assistance  
- **More enjoyable** terminal experience
- **Privacy-focused** with local AI models
- **Infinitely extensible** for your needs

---

## Roadmap

### Version 2.0 (Coming Soon)
- [ ] **Multi-Model Support**: GPT, Claude, and more AI providers
- [ ] **Project Templates**: Full application scaffolding
- [ ] **Git Integration**: AI-powered commit messages and PR reviews
- [ ] **Plugin System**: Custom command extensions
- [ ] **Web Dashboard**: Browser-based project management

### Version 3.0 (Future)
- [ ] **Team Collaboration**: Shared AI assistants
- [ ] **Cloud Sync**: Cross-device configuration sync
- [ ] **Advanced Analytics**: Development productivity insights
- [ ] **Voice Commands**: Speech-to-code capabilities
- [ ] **IDE Extensions**: Direct integration with popular editors

---

## License

MIT License - feel free to use DevLama in your projects, modify it, and distribute it.

---

## Support & Community

- **Issues**: [GitHub Issues](https://github.com/nitink23/devlama/issues)
- **Discussions**: [GitHub Discussions](https://github.com/nitink23/devlama/discussions)  
- **Documentation**: This README and [TESTING.md](TESTING.md)
- **Examples**: Check the `examples/` directory (coming soon)

---

## Acknowledgments

Built with love using:
- [Cobra](https://github.com/spf13/cobra) - Powerful CLI framework
- [OLLAMA](https://ollama.ai) - Local AI model management
- [PromptUI](https://github.com/manifoldco/promptui) - Interactive CLI components
- [Go](https://golang.org) - Fast, reliable systems programming

**DevLama: Where AI meets Developer Experience** 