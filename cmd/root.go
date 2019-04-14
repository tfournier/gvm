package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var version string

var rootCmd = &cobra.Command{
	Use:          "gvm",
	Short:        "Golang Version Manager",
	SilenceUsage: true,
	Version:      version,
}

// RootCmd export command
func RootCmd() *cobra.Command {
	return rootCmd
}

// Execute root command
func Execute() {
	if err := RootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
