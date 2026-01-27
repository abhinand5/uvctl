package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "uvctl",
	Short:         "Manage uv virtual environments",
	Long:          `uvctl is a CLI tool for managing Python virtual environments created using uv.`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(lsCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(activateCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(whichCmd)
	rootCmd.AddCommand(doctorCmd)
}
