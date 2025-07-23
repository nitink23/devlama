# DevLama - a CLI coding tool that BELONGS TO YOU

```
██████╗ ███████╗██╗   ██╗██╗      █████╗ ███╗   ███╗ █████╗ 
██╔══██╗██╔════╝██║   ██║██║     ██╔══██╗████╗ ████║██╔══██╗
██║  ██║█████╗  ██║   ██║██║     ███████║██╔████╔██║███████║
██║  ██║██╔══╝  ╚██╗ ██╔╝██║     ██╔══██║██║╚██╔╝██║██╔══██║
██████╔╝███████╗ ╚████╔╝ ███████╗██║  ██║██║ ╚═╝ ██║██║  ██║
╚═════╝ ╚══════╝  ╚═══╝  ╚══════╝╚═╝  ╚═╝╚═╝     ╚═╝╚═╝  ╚═╝
```

**A completely free, local alternative to Claude Code that works entirely on your machine**

DevLama brings AI-powered coding assistance directly to your terminal—no subscriptions, no data sharing, no limits. The quality of results depends entirely on your prompting skills and the local AI models you choose. Your code never leaves your machine.

---

## Why Choose DevLama Over Claude Code?

**DevLama gives you everything Claude Code offers, but completely free and under your control:**

- **100% Free** - No monthly subscriptions or usage limits
- **Completely Local** - Your code and data never leave your machine
- **Privacy First** - No data collection, tracking, or cloud dependency
- **Unlimited Usage** - Use as much as you want, whenever you want
- **Customizable Models** - Choose and configure your own AI models
- **Prompt-Dependent Quality** - Results improve with better prompting techniques

**Key Difference:** While Claude Code requires paid subscriptions and sends your code to external servers, DevLama runs entirely locally using OLLAMA, giving you complete control and privacy.

---

## What You Get (All Free)

### Everything Claude Code Does
- **Interactive AI Chat**: Real-time conversations with powerful local models
- **Code Analysis & Review**: AI-powered insights and suggestions
- **Smart Code Generation**: Create functions, classes, APIs with natural language
- **Bug Detection & Fixing**: Automatic identification and resolution of issues
- **Documentation Generation**: AI-generated comments and comprehensive docs
- **Multi-Language Support**: Go, Python, JavaScript, Java, C++, and more

### Plus Additional Advantages
- **No Cost**: Zero subscription fees, completely free forever
- **Local Processing**: All AI processing happens on your machine
- **Unlimited Usage**: No daily limits, token restrictions, or usage caps
- **Model Choice**: Use any OLLAMA-compatible model (llama2, codellama, mistral, etc.)
- **Full Privacy**: Your code never transmitted anywhere
- **Offline Capable**: Works without internet connection once models are downloaded

### Quality Depends on Your Skills
- **Prompt Engineering**: Better prompts = better results
- **Model Selection**: Choose the right model for your specific tasks
- **Context Provision**: More context in prompts leads to better outcomes
- **Iterative Refinement**: Learn to refine and improve your requests

---

## Quick Start

### Prerequisites
- Go 1.24.5+ - [Download Go](https://golang.org/dl/)
- OLLAMA - [Download OLLAMA](https://ollama.ai) - This replaces Claude Code's cloud dependency

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

### Setup Your Free AI Assistant (OLLAMA)

```bash
# 1. Install OLLAMA from https://ollama.ai (replaces Claude Code subscriptions)

# 2. Start OLLAMA service (runs locally, no internet required)
ollama serve

# 3. Download powerful AI models (one-time download, use forever)
ollama pull llama2        # General purpose model
ollama pull codellama     # Specialized for coding
ollama pull mistral       # Fast and efficient

# 4. Start using your free AI assistant
./devlama.exe ollama chat codellama
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

## DevLama vs Claude Code

### The Problem with Claude Code
- **Monthly subscriptions**: $20+ per month for full features
- **Usage limits**: Token restrictions and daily caps
- **Privacy concerns**: Your code is sent to external servers
- **Internet dependency**: Requires constant connection
- **No customization**: Locked into their models and interface

### DevLama's Solution
- **Completely free**: No subscriptions, no hidden costs
- **Unlimited usage**: Use as much as you want
- **Complete privacy**: Everything runs locally on your machine
- **Works offline**: Once models are downloaded, no internet needed
- **Full control**: Choose your models, customize everything

### Real-World Comparison
| Feature | Claude Code | DevLama |
|---------|-------------|---------|
| **Cost** | $20+/month | Free |
| **Privacy** | Code sent to servers | 100% local |
| **Usage Limits** | Yes | None |
| **Offline Mode** | No | Yes |
| **Model Choice** | Fixed | Any OLLAMA model |
| **Customization** | Limited | Complete control |
| **Data Control** | External | Your machine only |

### The Trade-off: Quality Depends on You
Unlike Claude Code's polished, consistent experience, DevLama's quality depends on:
- **Your prompting skills**: Learn to write effective prompts
- **Model selection**: Choose appropriate models for specific tasks
- **Context provision**: Give the AI enough information to work with
- **Iterative improvement**: Refine your approach over time

**Bottom Line**: DevLama gives you the same capabilities as Claude Code but free, private, and under your complete control.

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

**DevLama: Free AI coding assistance that respects your privacy and puts you in control** 