package cmd

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{"c"},
	Short:   "GVM configuration",
}

func init() {
	rootCmd.AddCommand(configCmd)
}
