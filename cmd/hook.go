package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var hookCmd = &cobra.Command{
	Use:   "hook <shell>",
	Short: "Print shell hook for seamless activation",
	Long: `Prints a shell hook that enables seamless 'uvctl activate' usage.

Add this to your shell configuration file:

  Bash (~/.bashrc):
    eval "$(uvctl hook bash)"

  Zsh (~/.zshrc):
    eval "$(uvctl hook zsh)"

After reloading your shell, you can simply run:
    uvctl activate <env-name>`,
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"bash", "zsh"},
	Run: func(cmd *cobra.Command, args []string) {
		shell := args[0]

		switch shell {
		case "bash", "zsh":
			printPosixHook()
		default:
			fmt.Fprintf(os.Stderr, "error: unsupported shell %q (supported: bash, zsh)\n", shell)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(hookCmd)
}

func printPosixHook() {
	// This creates a shell function that wraps uvctl
	// When 'activate' or 'deactivate' is called, it evals the output
	hook := `uvctl() {
    local cmd="${1:-}"

    if [ "$cmd" = "activate" ]; then
        if [ -z "${2:-}" ]; then
            command uvctl activate
            return $?
        fi
        eval "$(command uvctl activate "$2")"
    elif [ "$cmd" = "deactivate" ]; then
        eval "$(command uvctl deactivate)"
    else
        command uvctl "$@"
    fi
}`
	fmt.Println(hook)
}
