package cmd

import (
	"fmt"
	"os"

	"github.com/abhinand/uvctl/internal/config"
	"github.com/abhinand/uvctl/internal/uv"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Run diagnostics",
	Long:  `Performs environment diagnostics and reports any issues.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		hasErrors := false

		// Check uv availability
		fmt.Print("Checking uv installation... ")
		if uv.Available() {
			version, err := uv.Version()
			if err != nil {
				fmt.Printf("found, but cannot get version: %v\n", err)
			} else {
				fmt.Printf("ok (%s)\n", version)
			}

			uvPath, _ := uv.Path()
			fmt.Printf("  Path: %s\n", uvPath)
		} else {
			fmt.Println("NOT FOUND")
			fmt.Println("  Install uv: https://docs.astral.sh/uv/getting-started/installation/")
			hasErrors = true
		}

		fmt.Println()

		// Check root directory
		fmt.Print("Checking root directory... ")
		root, writable := config.IsRootWritable()
		if root == "" {
			fmt.Println("ERROR: cannot determine root directory")
			hasErrors = true
		} else if writable {
			fmt.Printf("ok (%s)\n", root)
		} else {
			fmt.Printf("NOT WRITABLE (%s)\n", root)
			fmt.Println("  Create the directory or set UVCTL_ROOT to a writable location")
			hasErrors = true
		}

		if os.Getenv(config.EnvRoot) != "" {
			fmt.Printf("  UVCTL_ROOT is set: %s\n", os.Getenv(config.EnvRoot))
		} else {
			fmt.Println("  UVCTL_ROOT is not set (using default)")
		}

		fmt.Println()

		// Check Python versions (only if uv is available)
		if uv.Available() {
			fmt.Println("Available Python versions:")
			pythons, err := uv.ListPythons()
			if err != nil {
				fmt.Printf("  Error listing Python versions: %v\n", err)
				hasErrors = true
			} else if len(pythons) == 0 {
				fmt.Println("  No Python versions found")
				fmt.Println("  Install Python via uv: uv python install 3.12")
			} else {
				// Show first few versions
				maxShow := 5
				for i, p := range pythons {
					if i >= maxShow {
						fmt.Printf("  ... and %d more\n", len(pythons)-maxShow)
						break
					}
					fmt.Printf("  %s\n", p)
				}
			}
		}

		fmt.Println()

		// Check active environment
		active := config.GetActive()
		if active != "" {
			fmt.Printf("Active environment: %s\n", active)
		} else {
			fmt.Println("No environment currently active")
		}

		if hasErrors {
			os.Exit(1)
		}
	},
}
