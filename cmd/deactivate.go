package cmd

import (
	"fmt"
	"os"

	"github.com/abhinand/uvctl/internal/config"
	"github.com/spf13/cobra"
)

var deactivateCmd = &cobra.Command{
	Use:   "deactivate",
	Short: "Deactivate the current environment",
	Long: `Deactivates the currently active uvctl environment.

Note: This command requires the shell hook to be installed.
Add this to your shell config:

    eval "$(uvctl hook bash)"   # or zsh

Then 'uvctl deactivate' will work seamlessly.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		active := config.GetActive()
		if active == "" {
			fmt.Fprintf(os.Stderr, "error: no uvctl environment is active\n")
			os.Exit(1)
		}

		// Print shell code to deactivate
		fmt.Println("# Deactivate current environment")
		fmt.Println("type deactivate &>/dev/null && deactivate")
		fmt.Println("unset UVCTL_ACTIVE")
	},
}

func init() {
	rootCmd.AddCommand(deactivateCmd)
}
