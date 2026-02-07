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

  Fish (~/.config/fish/config.fish):
    uvctl hook fish | source

After reloading your shell, you can simply run:
    uvctl activate <env-name>`,
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"bash", "zsh", "fish"},
	Run: func(cmd *cobra.Command, args []string) {
		shell := args[0]

		switch shell {
		case "bash", "zsh":
			printPosixHook()
		case "fish":
			printFishHook()
		default:
			fmt.Fprintf(os.Stderr, "error: unsupported shell %q (supported: bash, zsh, fish)\n", shell)
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

func printFishHook() {
	hook := `function uvctl
    set -l cmd $argv[1]

    if test "$cmd" = "activate"
        if test (count $argv) -lt 2
            command uvctl activate
            return $status
        end
        eval (command uvctl activate --shell fish $argv[2])
    else if test "$cmd" = "deactivate"
        eval (command uvctl deactivate --shell fish)
    else
        command uvctl $argv
    end
end`
	fmt.Println(hook)
}
