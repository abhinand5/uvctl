package cmd

import (
	"fmt"
	"os"

	"github.com/abhinand/uvctl/internal/env"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all environments",
	Long:  `Lists all available uvctl environments, one per line.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		envs, err := env.List()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		for _, name := range envs {
			fmt.Println(name)
		}
	},
}
