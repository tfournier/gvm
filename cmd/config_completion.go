package cmd

import (
	"github.com/spf13/cobra"
)

var configCompletionCmd = &cobra.Command{
	Use:     "completion",
	Aliases: []string{"c"},
	Short:   "Generates shell completion",
}

func init() {
	configCmd.AddCommand(configCompletionCmd)
}
