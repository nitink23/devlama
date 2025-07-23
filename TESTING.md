# DevLama CLI Testing Guide

This guide will help you test all the features of your DevLama CLI tool.

## Prerequisites

### For Basic Features
- Windows PowerShell (which you have)
- Built `devlama.exe` (already done)

### For OLLAMA Features
1. **Install OLLAMA**: Download from [ollama.ai](https://ollama.ai)
2. **Start OLLAMA service**: Run `ollama serve` in a separate terminal
3. **Download a model**: Run `ollama pull llama2` (or another model)

## Testing Steps

### 1. Basic CLI Tests

```bash
# Test the beautiful banner (default behavior)
.\devlama.exe

# Test help system
.\devlama.exe --help
.\devlama.exe help

# Test version command
.\devlama.exe version

# Test hello command
.\devlama.exe hello
.\devlama.exe hello --name "Nitin"
.\devlama.exe hello -n "Developer"

# Test open command (file/URL opener)
.\devlama.exe open .
.\devlama.exe open https://google.com
.\devlama.exe open README.md
```

### 2. Interactive Mode Tests

```bash
# Start interactive mode
.\devlama.exe interactive
# or use the alias:
.\devlama.exe i
```

**Once in interactive mode, try these commands:**
```
hello World
version
time
banner
echo Hello from DevLama!
open .
clear
help
exit
```

### 3. OLLAMA Integration Tests

⚠️ **Important**: Make sure OLLAMA is running first!

#### Step 3a: Start OLLAMA (in a separate PowerShell window)
```bash
# Download and start OLLAMA service
ollama serve
```

#### Step 3b: Download a Model (in another PowerShell window)
```bash
# Download a lightweight model for testing
ollama pull llama2
# or try a smaller model:
ollama pull tinyllama
```

#### Step 3c: Test OLLAMA Commands
```bash
# List available models
.\devlama.exe ollama list

# Test one-time text generation
.\devlama.exe ollama generate llama2 "Write a haiku about programming"

# Test interactive chat
.\devlama.exe ollama chat llama2

# Test model pulling
.\devlama.exe ollama pull tinyllama
```

#### Step 3d: Test OLLAMA in Interactive Mode
```bash
# Start interactive mode
.\devlama.exe interactive

# Then try these commands:
ollama list
ollama chat llama2
ollama generate llama2 Tell me a joke
exit
```

### 4. Advanced Testing Scenarios

#### Test Error Handling
```bash
# Test with OLLAMA not running
.\devlama.exe ollama list

# Test with non-existent model
.\devlama.exe ollama chat nonexistent-model

# Test invalid commands
.\devlama.exe invalid-command
```

#### Test File Operations
```bash
# Test opening different file types
.\devlama.exe open go.mod
.\devlama.exe open cmd/
.\devlama.exe open https://github.com
```

## Expected Results

### ✅ Successful Tests Should Show:

1. **Banner Display**: Colorful ASCII art with system info
2. **Interactive Mode**: Responsive text input box with colored prompts
3. **OLLAMA Chat**: Real-time AI responses streaming in
4. **File Opening**: Files/URLs opening in default applications
5. **Help System**: Well-formatted command documentation

### 🔍 What to Look For:

- **Colors and Emojis**: Commands should display with colors and emoji icons
- **Error Messages**: Clear, helpful error messages with suggestions
- **Progress Indicators**: Progress bars when downloading models
- **Responsive UI**: Smooth text input and streaming responses

## Troubleshooting

### OLLAMA Issues
```bash
# If OLLAMA commands fail:
1. Check if OLLAMA is running: ollama serve
2. Check if models are available: ollama list
3. Try downloading a model: ollama pull llama2
4. Check OLLAMA status: ollama ps
```

### General Issues
```bash
# If CLI doesn't work:
1. Rebuild: go build -o devlama.exe
2. Check permissions: Make sure devlama.exe can execute
3. Check dependencies: go mod tidy
```

## Test Script

Here's a quick test script you can run:

```powershell
# Save as test-devlama.ps1
Write-Host "🧪 Testing DevLama CLI..." -ForegroundColor Cyan

Write-Host "`n1. Testing banner..." -ForegroundColor Yellow
.\devlama.exe

Write-Host "`n2. Testing version..." -ForegroundColor Yellow
.\devlama.exe version

Write-Host "`n3. Testing hello..." -ForegroundColor Yellow
.\devlama.exe hello --name "Tester"

Write-Host "`n4. Testing help..." -ForegroundColor Yellow
.\devlama.exe --help

Write-Host "`n5. Testing OLLAMA (if available)..." -ForegroundColor Yellow
.\devlama.exe ollama list

Write-Host "`n✅ Basic tests complete!" -ForegroundColor Green
Write-Host "💡 Try interactive mode with: .\devlama.exe interactive" -ForegroundColor Cyan
```

## Performance Expectations

- **Banner Display**: Instant
- **File Operations**: < 1 second
- **OLLAMA Model Download**: 1-5 minutes (depending on model size)
- **AI Chat Responses**: 2-10 seconds per response
- **Interactive Mode**: Real-time responsiveness

## Next Steps

After testing, you can:
1. **Customize**: Modify models, prompts, or commands
2. **Extend**: Add new OLLAMA models or features
3. **Deploy**: Share the `devlama.exe` with others
4. **Automate**: Create scripts for common workflows

Happy testing! 🎉 