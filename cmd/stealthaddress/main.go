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
    // Flag to enable verbose logging.
    verbose := flag.Bool("verbose", false, "Enable verbose logging")
    // Input file path for the application.
    input := flag.String("input", "", "Input file path")
    // Output file path for the application.
    output := flag.String("output", "", "Output file path")
    flag.Parse()

    // Create a new StealthAddress application instance with the provided verbose flag.
    // This instance will be used to run the application.
    app := stealthaddress.NewApp(*verbose)

    // Run the application with the provided input and output file paths.
    // If an error occurs during execution, it will be logged and the application will exit with a non-zero status code.
    if err := app.Run(*input, *output); err != nil {
        log.Printf("Error: %v", err)
        os.Exit(1)
    }
}