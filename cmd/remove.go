package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tfournier/gvm/src"
)

var removeCmd = &cobra.Command{
	Use:     "remove [version]",
	Aliases: []string{"r"},
	Short:   "Remove SDK version",
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		purge, err := cmd.Flags().GetBool("purge")
		if err != nil {
			return err
		}
		return src.GVM().SDK().Uninstall(args[0], purge)
	},
}

func init() {
	removeCmd.Flags().Bool("purge", false, "Remove downloaded archive")
	rootCmd.AddCommand(removeCmd)
}
