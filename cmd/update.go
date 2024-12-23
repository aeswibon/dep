package cmd

import (
	"fmt"

	"github.com/aeswibon/dep/pkg/manager"
	"github.com/aeswibon/dep/pkg/utils"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update [dependencies]",
	Short:   "Update dependencies",
	Long:    "Update dependencies using the specified package manager.",
	GroupID: "dependenciesGroup",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Load the configuration file
		config, err := utils.LoadConfig()
		if err != nil {
			fmt.Printf("Error loading configuration: %v\n", err)
			return
		}

		// Get the package manager from the configuration
		pm, err := manager.GetPackageManager(config.PackageManager)
		if err != nil {
			fmt.Printf("Error getting package manager: %v\n", err)
			return
		}

		// check if the command is available
		if err := utils.CheckTool(pm.GetCommand()); err != nil {
			fmt.Printf("Error checking command: %v\n", err)
			return
		}

		// Execute the add command
		if err := pm.ExecuteCommand("update", args, config, &globalFlags); err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			return
		}
	},
}

var upgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Short:   "Upgrade all dependencies",
	Long:    "Upgrade all dependencies using the specified package manager.",
	GroupID: "dependenciesGroup",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Load the configuration file
		config, err := utils.LoadConfig()
		if err != nil {
			fmt.Printf("Error loading configuration: %v\n", err)
			return
		}

		// Get the package manager from the configuration
		pm, err := manager.GetPackageManager(config.PackageManager)
		if err != nil {
			fmt.Printf("Error getting package manager: %v\n", err)
			return
		}

		// check if the command is available
		if err := utils.CheckTool(pm.GetCommand()); err != nil {
			fmt.Printf("Error checking command: %v\n", err)
			return
		}

		// Execute the add command
		if err := pm.ExecuteCommand("upgrade", args, config, &globalFlags); err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(upgradeCmd)
}
