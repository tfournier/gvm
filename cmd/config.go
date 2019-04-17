package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tfournier/gvm/src"
)

var configCmd = &cobra.Command{
	Use:     "config [command]",
	Aliases: []string{"c"},
	Short:   "GVM configuration",
	Args:    cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		src.GVM().ShowConfig()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
