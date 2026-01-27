package cmd

import (
	"fmt"
	"os"

	"github.com/abhinand/uvctl/internal/env"
	"github.com/abhinand/uvctl/internal/uv"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create <env-name> <python-version>",
	Short: "Create a new environment",
	Long:  `Creates a new uv virtual environment with the specified Python version.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		pythonVersion := args[1]

		// Check uv is available first
		if !uv.Available() {
			fmt.Fprintf(os.Stderr, "error: uv not found in PATH\n")
			os.Exit(1)
		}

		// Create the environment
		if err := env.Create(name, pythonVersion); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	},
}
