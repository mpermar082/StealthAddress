// cmd/stealthaddress/main.go
package main

import (
    "flag"
    "log"
    "os"

    "stealthaddress/internal/stealthaddress"
)

// Main entry point for the StealthAddress application.
func main() {
    // Define command-line flags for the application.
    verbose := flag.Bool("verbose", false, "Enable verbose logging")
    input := flag.String("input", "", "Input file path")
    output := flag.String("output", "", "Output file path")
    flag.Parse()

    // Create a new StealthAddress application instance with the provided verbose flag.
    app := stealthaddress.NewApp(*verbose)

    // Run the application with the provided input and output file paths.
    if err := app.Run(*input, *output); err != nil {
        // Log any errors that occur during application execution.
        log.Printf("Error: %v", err)
        // Exit the application with a non-zero status code to indicate failure.
        os.Exit(1)
    }
}