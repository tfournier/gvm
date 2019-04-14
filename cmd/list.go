package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tfournier/gvm/src"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List Golang version",
	RunE: func(cmd *cobra.Command, args []string) error {
		return src.GVM().SDK().ShowList()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
