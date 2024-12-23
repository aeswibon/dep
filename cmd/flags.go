package cmd

import (
	"github.com/aeswibon/dep/pkg/types"
	"github.com/spf13/cobra"
)

var globalFlags types.GlobalFlags

// Init initializes the global flags
func Init(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().BoolVarP(&globalFlags.Verbose, "verbose", "v", false, "enable verbose output")
	rootCmd.PersistentFlags().StringVarP(&globalFlags.ConfigFile, "config", "c", "", "config file (default is config.json)")
}
