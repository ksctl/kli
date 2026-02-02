package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func (k *KsctlCommand) ShellCompletion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "completion [bash|zsh|fish]",
		Short: "Generate shell completion scripts",
		Long: `To load completions:

Bash:

  $ source <(kli completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ kli completion bash > /etc/bash_completion.d/kli
  # macOS:
  $ kli completion bash > /usr/local/etc/bash_completion.d/kli

Zsh:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc
  $ kli completion zsh > "${fpath[1]}/_kli"

Fish:

  $ kli completion fish | source

  # To load completions for each session, execute once:
  $ kli completion fish > ~/.config/fish/completions/kli.fish
`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			switch args[0] {
			case "bash":
				return cmd.Root().GenBashCompletion(os.Stdout)
			case "zsh":
				return cmd.Root().GenZshCompletion(os.Stdout)
			case "fish":
				return cmd.Root().GenFishCompletion(os.Stdout, true)
			default:
				return fmt.Errorf("unsupported shell: %s", args[0])
			}
		},
	}

	return cmd
}
