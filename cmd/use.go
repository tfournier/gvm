package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tfournier/gvm/src"
)

var useCmd = &cobra.Command{
	Use:     "use [version]",
	Aliases: []string{"u"},
	Short:   "Switch SDK",
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return src.GVM().SDK().Switch(args[0])
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
