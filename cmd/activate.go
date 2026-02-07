package cmd

import (
	"fmt"
	"os"

	"github.com/abhinand5/uvctl/internal/env"
	"github.com/spf13/cobra"
)

var activateShellFlag string

var activateCmd = &cobra.Command{
	Use:   "activate <env-name>",
	Short: "Print shell commands to activate an environment",
	Long: `Prints shell commands to stdout that will activate the specified environment.

Usage:
  eval "$(uvctl activate <env-name>)"`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		// Validate the environment exists and is ready
		if err := env.ValidateActivate(name); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		if activateShellFlag == "fish" {
			activatePath, err := env.FishActivatePath(name)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}

			fmt.Println("# Deactivate any existing virtualenv")
			fmt.Println("functions -q deactivate; and deactivate")
			fmt.Println("")
			fmt.Println("# Activate the new environment")
			fmt.Printf("source %q\n", activatePath)
			fmt.Println("")
			fmt.Println("# Set uvctl tracking variable")
			fmt.Printf("set -gx UVCTL_ACTIVE %q\n", name)
		} else {
			activatePath, err := env.ActivatePath(name)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}

			// Print shell code to stdout
			// This must be eval'd by the user: eval "$(uvctl activate <env>)"
			fmt.Println("# Deactivate any existing virtualenv")
			fmt.Println("type deactivate &>/dev/null && deactivate")
			fmt.Println("")
			fmt.Println("# Activate the new environment")
			fmt.Printf("source %q\n", activatePath)
			fmt.Println("")
			fmt.Println("# Set uvctl tracking variable")
			fmt.Printf("export UVCTL_ACTIVE=%q\n", name)
		}
	},
}

func init() {
	activateCmd.Flags().StringVar(&activateShellFlag, "shell", "", "shell type for output format (fish)")
}
