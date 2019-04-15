package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tfournier/zshcompletion"
)

var configCompletionZshCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Generates shell completion for ZSH",
	RunE: func(cmd *cobra.Command, args []string) error {
		return zshcompletion.Zsh(cmd.Root()).GenCompletionUser()
	},
}

func init() {
	configCompletionCmd.AddCommand(configCompletionZshCmd)
}
