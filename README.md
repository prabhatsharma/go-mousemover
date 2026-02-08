# go-mousemover

A simple, modern utility that moves your mouse periodically to prevent automatic screen lock.

## Overview

This tool moves your mouse by a configurable number of pixels at regular intervals. It's useful when you have automatic screen lock enabled on your computer and you're reading or thinking without actively using the mouse/keyboard, preventing annoying interruptions.

**⚠️ Warning**: Running this tool may violate your company's IT policies. Use at your own risk!

## Features

- ✨ **Configurable interval** - Set custom time between mouse movements
- 🎯 **Configurable distance** - Set custom pixel distance for movements
- 🔄 **Graceful shutdown** - Clean exit with Ctrl+C (SIGINT/SIGTERM handling)
- 📊 **Verbose mode** - Optional detailed logging of mouse positions
- 🚀 **Modern Go code** - Uses contexts, goroutines, and proper signal handling
- 🏗️ **Cross-platform** - Works on macOS, Linux, and Windows
- 📦 **Go modules** - Proper dependency management

## Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/prabhatsharma/go-mousemover.git
cd go-mousemover

# Build the binary
make build

# Or install to GOPATH/bin
make install
```

### Using Go Install

```bash
go install github.com/prabhatsharma/go-mousemover@latest
```

## Usage

### Basic Usage

Run with default settings (moves mouse by 1 pixel every 30 seconds):

```bash
go-mousemover
```

or with the Makefile:

```bash
make run
```

### Command-Line Options

```bash
go-mousemover [options]

Options:
  -interval int     Interval in seconds between mouse movements (default: 30)
  -distance int     Distance in pixels to move the mouse (default: 1)
  -verbose          Enable verbose logging
  -v                Enable verbose logging (shorthand)
  -version          Print version information
```

### Examples

```bash
# Move mouse every 60 seconds
go-mousemover -interval 60

# Move mouse by 5 pixels every 30 seconds with verbose logging
go-mousemover -distance 5 -verbose

# Move mouse every 2 minutes (120 seconds) by 10 pixels
go-mousemover -interval 120 -distance 10

# Show version
go-mousemover -version
```

## Building

### Build for Current Platform

```bash
make build
```

The binary will be created in the `bin/` directory.

### Build for All Platforms

```bash
make build-all
```

This creates binaries for:
- Linux (amd64)
- macOS (amd64, arm64)
- Windows (amd64)

### Available Make Targets

```bash
make help              # Show all available commands
make build             # Build the application
make run               # Run the application
make run-verbose       # Run with verbose logging
make clean             # Clean build artifacts
make install           # Install to GOPATH/bin
make test              # Run tests
make fmt               # Format Go code
make deps              # Download dependencies
make lint              # Run linter (requires golangci-lint)
make build-all         # Build for all platforms
make version           # Show version
```

## Development

### Requirements

- Go 1.16 or higher
- Make (optional, for using Makefile)

### Project Structure

```
go-mousemover/
├── main.go          # Main application code
├── go.mod           # Go module definition
├── go.sum           # Dependency checksums
├── Makefile         # Build automation
├── README.md        # This file
├── LICENSE          # License file
└── .gitignore       # Git ignore rules
```

### Dependencies

- [robotgo](https://github.com/go-vgo/robotgo) - Cross-platform Go library for controlling mouse and keyboard

## How It Works

1. The application starts with your configured settings
2. It captures the current mouse position
3. Moves the mouse by the specified distance (both X and Y coordinates)
4. Waits for the specified interval
5. Repeats until you press Ctrl+C

The tool uses goroutines and context for proper concurrent execution and graceful shutdown handling.

## Stopping the Application

Press `Ctrl+C` to stop the application. It will shut down gracefully and display a shutdown message.

## License

See the [LICENSE](LICENSE) file for details.

## Author

Created by **Prabhat Sharma** (<hi.prabhat@gmail.com>)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Disclaimer

This tool is provided as-is for personal use. Running utilities that simulate user activity may violate workplace policies or terms of service. Always ensure you have permission to use such tools in your environment.
