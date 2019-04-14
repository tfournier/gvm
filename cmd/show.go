package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tfournier/gvm/src"
)

var showCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"s"},
	Short:   "Show SDK",
	Args: func(cmd *cobra.Command, args []string) error {
		cobra.MinimumNArgs(1)
		if len(args) == 0 {
			return fmt.Errorf("version required")
		}
		if !src.GVM().SDK().HasValidVersion(args[0]) {
			return fmt.Errorf("not valid version")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		src.GVM().SDK().Show(args[0])
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
