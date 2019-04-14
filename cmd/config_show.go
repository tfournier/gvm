package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tfournier/gvm/src"
)

var configShowCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"s"},
	Short:   "Display environment configuration",
	Run: func(cmd *cobra.Command, args []string) {
		src.GVM().ShowConfig()
	},
}

func init() {
	configCmd.AddCommand(configShowCmd)
}
