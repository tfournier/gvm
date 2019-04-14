package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tfournier/gvm/src"
)

var installCmd = &cobra.Command{
	Use:     "install <version>",
	Aliases: []string{"i"},
	Short:   "Install SDK",
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return src.GVM().SDK().Install(args[0])
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
