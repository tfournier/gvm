package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tfournier/gvm/src"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Info SDK used",
	RunE: func(cmd *cobra.Command, args []string) error {
		return src.GVM().SDK().Info()
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
