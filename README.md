# DevLama - a CLI coding tool that BELONGS TO YOU

> **⚠️ EXPERIMENTAL PROJECT - NOT READY FOR PRODUCTION USE**
> 
> DevLama is currently in **heavy development** and **does not work yet**. Many features are incomplete or non-functional. This is a proof-of-concept and demonstration of what the tool aims to become. 
>
> **Current Status**: Basic CLI structure works, but AI integration and code manipulation features are still being developed.
>
> **Expected Timeline**: Alpha version targeting Q2 2024

```
██████╗ ███████╗██╗   ██╗██╗      █████╗ ███╗   ███╗ █████╗ 
██╔══██╗██╔════╝██║   ██║██║     ██╔══██╗████╗ ████║██╔══██╗
██║  ██║█████╗  ██║   ██║██║     ███████║██╔████╔██║███████║
██║  ██║██╔══╝  ╚██╗ ██╔╝██║     ██╔══██║██║╚██╔╝██║██╔══██║
██████╔╝███████╗ ╚████╔╝ ███████╗██║  ██║██║ ╚═╝ ██║██║  ██║
╚═════╝ ╚══════╝  ╚═══╝  ╚══════╝╚═╝  ╚═╝╚═╝     ╚═╝╚═╝  ╚═╝
```

**A completely free, local alternative to Claude Code that works entirely on your machine**

DevLama aims to bring AI-powered coding assistance directly to your terminal—no subscriptions, no data sharing, no limits. The quality of results will depend entirely on your prompting skills and the local AI models you choose. Your code never leaves your machine.

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

## Development Status & Quick Start

> **⚠️ Warning**: These installation instructions are for **development and testing only**. Most features don't work yet.

### Prerequisites
- Go 1.24.5+ - [Download Go](https://golang.org/dl/)
- OLLAMA - [Download OLLAMA](https://ollama.ai) - This will replace Claude Code's cloud dependency

### Installation (For Developers)

```bash
# Clone the repository
git clone https://github.com/nitink23/devlama.git
cd devlama

# Build DevLama
go build -o devlama.exe  # Windows
# go build -o devlama     # macOS/Linux

# Test basic functionality (limited features work)
./devlama.exe --help
./devlama.exe version
```

### Setup Your Free AI Assistant (OLLAMA) - **EXPERIMENTAL**

> **Note**: AI features are still in development and may not work properly.

```bash
# 1. Install OLLAMA from https://ollama.ai (will replace Claude Code subscriptions)

# 2. Start OLLAMA service (runs locally, no internet required)
ollama serve

# 3. Download powerful AI models (one-time download, use forever)
ollama pull llama2        # General purpose model
ollama pull codellama     # Specialized for coding
ollama pull mistral       # Fast and efficient

# 4. Try the experimental AI features (may not work)
./devlama.exe ollama chat codellama
```

---

## Usage Examples (Planned Features)

> **⚠️ Important**: Most of these features are **not implemented yet**. This shows what DevLama will be capable of when development is complete.

### Interactive Mode (Planned)
```bash
# Start the beautiful interactive interface (basic version works)
./devlama.exe interactive

# Try these commands in the text box (limited functionality):
> hello DevLama        # ✅ Works
> banner               # ✅ Works  
> version              # ✅ Works
> ollama chat llama2   # ⚠️ Experimental
> code edit main.go    # ❌ Not implemented
> code generate        # ❌ Not implemented
> help
```

### AI Commands (Experimental)
```bash
# Chat with AI models (may not work properly)
./devlama.exe ollama chat llama2      # ⚠️ Experimental
./devlama.exe ollama list             # ⚠️ Experimental  
./devlama.exe ollama pull codellama   # ⚠️ Experimental

# Generate text with AI (not fully implemented)
./devlama.exe ollama generate llama2 "Write a Go function to reverse a string"  # ❌ May fail
```

### Code Manipulation (Not Implemented)
```bash
# Edit code files with AI assistance (planned feature)
./devlama.exe code edit main.go                              # ❌ Not working yet

# Generate new code from templates (planned feature)
./devlama.exe code generate function                         # ❌ Not working yet
./devlama.exe code generate api                              # ❌ Not working yet
./devlama.exe code generate test                             # ❌ Not working yet

# Analyze and improve existing code (planned feature)
./devlama.exe code analyze cmd/root.go                       # ❌ Not working yet
./devlama.exe code refactor utils.go                         # ❌ Not working yet

# Add features to existing files (planned feature)
./devlama.exe code add "user authentication system"          # ❌ Not working yet
```

### Utility Commands (Working)
```bash
# Open files and URLs (works)
./devlama.exe open .                    # ✅ Works
./devlama.exe open https://github.com   # ✅ Works

# Show system information (works)
./devlama.exe version                   # ✅ Works
./devlama.exe                          # ✅ Works (shows banner)

# Interactive help (works)
./devlama.exe help                     # ✅ Works
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

## Testing Guide (For Developers)

We've created a comprehensive testing guide! See **[TESTING.md](TESTING.md)** for:
- Step-by-step testing instructions
- OLLAMA setup verification
- Feature validation checklist (shows what works vs. what doesn't)
- Troubleshooting common issues
- Performance benchmarks

### Quick Test (What Actually Works)
```bash
# Test basic functionality (should work)
./devlama.exe version                 # ✅ Works
./devlama.exe hello --name "Tester"   # ✅ Works

# Test AI integration (experimental, may fail)
./devlama.exe ollama list             # ⚠️ May not work

# Test interactive mode (basic version works)
./devlama.exe interactive             # ✅ Basic functionality works
```

---

## Contributing (Help Wanted!)

**DevLama needs your help!** Since this is an experimental project under heavy development, contributions are especially welcome. Here's how to get started:

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

### High Priority Contribution Areas
- **AI Integration**: Complete the OLLAMA implementation (partially working)
- **Code Manipulation**: Implement the code editing and generation features
- **Error Handling**: Improve error messages and edge case handling
- **Testing**: Add comprehensive tests for existing functionality
- **Documentation**: Update docs as features become functional
- **Bug Fixes**: Many known issues need fixing

### Current Development Status
- ✅ **Basic CLI structure** - Working
- ✅ **Interactive mode** - Basic version working
- ✅ **File operations** - Working
- ⚠️ **OLLAMA integration** - Partially implemented
- ❌ **Code editing** - Not implemented
- ❌ **AI code generation** - Not implemented
- ❌ **Code analysis** - Not implemented

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