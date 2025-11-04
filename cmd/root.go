// Package cmd provides the command-line interface implementation for GoLoad,
// a HTTP load testing tool. It contains all command definitions, argument parsing,
// and execution handlers for the goload CLI application.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// base command
var rootCmd = &cobra.Command{
	Use:   "goload",
	Short: "GoLoad - a HTTP Load Testing tool written in Go.",
	Long: `GoLoad is a lightweight HTTP load testing tool written in Go.

Designed for developers and DevOps engineers who need to quickly assess
the performance and reliability of HTTP services under various load conditions.

Aimed to be perfect for testing APIs, web services, and microservices in development,
staging, or production environments.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "There was an error running the root command: %v", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
