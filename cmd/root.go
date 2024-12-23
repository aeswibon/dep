package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var verbose bool

var rootCmd = &cobra.Command{
	Use:   "dep",
	Short: "A CLI tool to manage dependencies",
	Long:  `Dep is a CLI tool to manage dependencies for different package managers, such as Go modules, npm, and pip. It can also run custom commands defined in the configuration file.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if verbose {
			fmt.Println("Running in verbose mode")
		}
	},
}

// Execute runs the root command
func Execute() {
	rootCmd.AddGroup(dependenciesGroup)
	rootCmd.AddGroup(scriptGroup)
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
}
