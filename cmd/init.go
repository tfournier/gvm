package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tfournier/gvm/src"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Initialize GVM",
	Args:    cobra.MaximumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return src.GVM().Initialize(cmd.Root())
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
