# Greeter CLI

A simple yet fun command-line greeting application built with Go. Greet users in multiple languages, get time-based greetings, and enjoy interactive features.

## Features

- **Multi-language support**: Greet in English, Spanish, French, Swahili, German, and Italian
- **Time-based greetings**: Get appropriate greetings based on the time of day
- **Interactive mode**: Ask for user input and personalize greetings
- **Colored output**: Enjoy colorful terminal output
- **Configuration**: Save and load default names
- **Fun random messages**: Get motivational quotes

## Installation

```bash
go build -o greeter
```

Or run directly:

```bash
go run main.go <command> [flags]
```

## Usage

### Commands

| Command | Description | Example |
|---------|-------------|---------|
| `hello` | Greet someone | `go run main.go hello --name Brian` |
| `bye` | Say goodbye | `go run main.go bye --name Brian` |
| `time` | Time-based greeting | `go run main.go time` |
| `lang` | Multi-language greeting | `go run main.go lang --lang es --name Brian` |
| `random` | Random motivational message | `go run main.go random` |
| `ask` | Interactive greeting | `go run main.go ask` |
| `version` | Show version info | `go run main.go version` |

### Examples

```bash
# Simple greeting
go run main.go hello --name Brian

# Loud greeting
go run main.go hello --name Brian --loud

# Goodbye with colors
go run main.go bye --name Brian

# Time-aware greeting
go run main.go time

# Greet in Spanish
go run main.go lang --lang es --name Brian

# Get a random motivational message
go run main.go random

# Interactive mode
go run main.go ask
```

### Supported Languages

- `en` - English
- `es` - Spanish
- `fr` - French
- `de` - German
- `sw` - Swahili
- `it` - Italian

## Configuration

Save your default name:

```bash
go run main.go hello --name Brian
```

Your default name is stored in `~/.greeterconfig` and will be used if no name is provided.

## Project Structure

```
go-greeter/
├── cmd/               # Command implementations
│   ├── hello.go       # Hello command
│   ├── bye.go         # Bye command
│   ├── time.go        # Time-based greeting
│   ├── lang.go        # Multi-language support
│   ├── random.go      # Random messages
│   ├── interactive.go # Interactive mode
│   ├── help.go        # Help command
│   └── root.go        # Root command setup
├── utils/             # Utility functions
│   ├── colors.go      # ANSI color codes
│   ├── config.go      # Configuration management
│   ├── lang.go        # Language definitions
│   └── banner.go      # CLI banner display
├── main.go            # Application entry point
├── go.mod             # Go module definition
└── README.md          # This file
```

## Version

Greeter CLI v3.0.0