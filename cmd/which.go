package cmd

import (
	"fmt"
	"os"

	"github.com/abhinand5/uvctl/internal/config"
	"github.com/abhinand5/uvctl/internal/env"
	"github.com/spf13/cobra"
)

var whichCmd = &cobra.Command{
	Use:   "which",
	Short: "Print the active environment path",
	Long:  `Prints the absolute path of the currently active uvctl environment.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		active := config.GetActive()
		if active == "" {
			fmt.Fprintf(os.Stderr, "error: no environment is currently active\n")
			os.Exit(1)
		}

		// Verify the environment still exists
		exists, err := env.Exists(active)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		if !exists {
			fmt.Fprintf(os.Stderr, "error: active environment %q no longer exists\n", active)
			os.Exit(1)
		}

		path, err := env.Path(active)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(path)
	},
}
