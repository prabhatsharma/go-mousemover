package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-vgo/robotgo"
)

const (
	appName    = "go-mousemover"
	appVersion = "2.0.0"
	appAuthor  = "Prabhat Sharma <hi.prabhat@gmail.com>"
)

// Config holds application configuration
type Config struct {
	Interval int
	Distance int
	Verbose  bool
	Version  bool
}

func main() {
	config := parseFlags()

	if config.Version {
		printVersion()
		os.Exit(0)
	}

	printBanner(config)

	// Create context with cancellation for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Setup signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Start mouse mover in a goroutine
	done := make(chan bool)
	go runMouseMover(ctx, config, done)

	// Wait for interrupt signal
	<-sigChan
	log.Println("\nShutting down gracefully...")
	cancel()

	// Wait for mouse mover to finish
	<-done
	log.Println("Shutdown complete")
}

// parseFlags parses command-line flags
func parseFlags() *Config {
	config := &Config{}

	flag.IntVar(&config.Interval, "interval", 30, "Interval in seconds between mouse movements")
	flag.IntVar(&config.Distance, "distance", 1, "Distance in pixels to move the mouse")
	flag.BoolVar(&config.Verbose, "verbose", false, "Enable verbose logging")
	flag.BoolVar(&config.Verbose, "v", false, "Enable verbose logging (shorthand)")
	flag.BoolVar(&config.Version, "version", false, "Print version information")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", appName)
		fmt.Fprintf(os.Stderr, "A simple utility that moves your mouse periodically to prevent screen lock.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  %s                           # Default: 30 second interval, 1 pixel movement\n", appName)
		fmt.Fprintf(os.Stderr, "  %s -interval 60              # Move mouse every 60 seconds\n", appName)
		fmt.Fprintf(os.Stderr, "  %s -distance 5 -v            # Move 5 pixels with verbose logging\n", appName)
		fmt.Fprintf(os.Stderr, "\nWarning:\n")
		fmt.Fprintf(os.Stderr, "  Running this tool may violate your company's IT policies.\n")
		fmt.Fprintf(os.Stderr, "  Use at your own risk!\n")
	}

	flag.Parse()

	// Validate configuration
	if config.Interval < 1 {
		log.Fatal("Interval must be at least 1 second")
	}
	if config.Distance < 1 {
		log.Fatal("Distance must be at least 1 pixel")
	}

	return config
}

// printBanner prints the application banner
func printBanner(config *Config) {
	fmt.Printf("=== %s v%s ===\n", appName, appVersion)
	fmt.Printf("Created by %s\n", appAuthor)
	fmt.Printf("Configuration:\n")
	fmt.Printf("  - Interval: %d seconds\n", config.Interval)
	fmt.Printf("  - Distance: %d pixels\n", config.Distance)
	fmt.Printf("  - Verbose:  %v\n", config.Verbose)
	fmt.Println("\nPress Ctrl+C to stop")
	fmt.Println("=======================")
	fmt.Println()
}

// printVersion prints version information
func printVersion() {
	fmt.Printf("%s version %s\n", appName, appVersion)
	fmt.Printf("Created by %s\n", appAuthor)
}

// runMouseMover runs the main mouse movement loop
func runMouseMover(ctx context.Context, config *Config, done chan bool) {
	defer close(done)

	ticker := time.NewTicker(time.Duration(config.Interval) * time.Second)
	defer ticker.Stop()

	moveCount := 0

	for {
		select {
		case <-ctx.Done():
			if config.Verbose {
				log.Printf("Context cancelled, stopping mouse mover (total moves: %d)", moveCount)
			}
			return
		case <-ticker.C:
			if err := moveMouse(config); err != nil {
				log.Printf("Error moving mouse: %v", err)
				continue
			}
			moveCount++

			if config.Verbose {
				x, y := robotgo.Location()
				log.Printf("Move #%d - Mouse position: x=%d, y=%d", moveCount, x, y)
			} else {
				fmt.Print(".")
			}
		}
	}
}

// moveMouse moves the mouse by the specified distance
func moveMouse(config *Config) error {
	x, y := robotgo.Location()

	if config.Verbose {
		log.Printf("Current position: x=%d, y=%d", x, y)
	}

	newX := x + config.Distance
	newY := y + config.Distance

	robotgo.Move(newX, newY)

	if config.Verbose {
		log.Printf("Moved to: x=%d, y=%d", newX, newY)
	}

	return nil
}
