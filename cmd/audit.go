package cmd

import (
	"fmt"

	"github.com/aeswibon/dep/pkg/manager"
	"github.com/aeswibon/dep/pkg/utils"
	"github.com/spf13/cobra"
)

var auditCmd = &cobra.Command{
	Use:     "audit",
	Short:   "Audit the dependencies",
	Long:    "Audit the dependencies of the package manager.",
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
		if err := pm.ExecuteCommand("audit", args, config, &globalFlags); err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			return
		}
	},
}

var fixVulnerabilitiesCmd = &cobra.Command{
	Use:     "fix-audit",
	Short:   "Fix the vulnerabilities",
	Long:    "Fix the vulnerabilities of the package manager.",
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
		if err := pm.ExecuteCommand("fix-audit", args, config, &globalFlags); err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(auditCmd)
	rootCmd.AddCommand(fixVulnerabilitiesCmd)
}
