package cmd

import (
	"fmt"
	"os"

	"github.com/abhinand/uvctl/internal/env"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <env-name>",
	Short: "Delete an environment",
	Long:  `Deletes an environment and all its contents. This operation is non-interactive.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		if err := env.Delete(name); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	},
}
