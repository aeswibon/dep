package cmd

import (
	"fmt"

	"github.com/aeswibon/dep/pkg/manager"
	"github.com/aeswibon/dep/pkg/utils"
	"github.com/spf13/cobra"
)

// global commands

var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Install global dependencies",
	Long:    "Install global dependencies using the specified package manager.",
	GroupID: "globalGroup",
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
		if err := pm.ExecuteCommand("install", args, config, &globalFlags); err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			return
		}
	},
}

var uninstallCmd = &cobra.Command{
	Use:     "uninstall [dependencies]",
	Short:   "Remove global dependencies",
	Long:    "Remove global dependencies using the specified package manager.",
	GroupID: "globalGroup",
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
		if err := pm.ExecuteCommand("uninstall", args, config, &globalFlags); err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			return
		}
	},
}

// install commands
var addCmd = &cobra.Command{
	Use:     "add [dependencies]",
	Short:   "Add dependencies to the project",
	Long:    "Add dependencies to the project using the specified package manager.",
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
		if err := pm.ExecuteCommand("add", args, config, &globalFlags); err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			return
		}
	},
}

var localCmd = &cobra.Command{
	Use:     "local",
	Short:   "Install all the dependencies",
	Long:    "Install all the dependencies to the project using the specified package manager.",
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
		if err := pm.ExecuteCommand("local", args, config, &globalFlags); err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			return
		}
	},
}

var devCmd = &cobra.Command{
	Use:     "dev [dependencies]",
	Short:   "Install dev dependencies",
	Long:    "Install dev dependencies using the specified package manager.",
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
		if err := pm.ExecuteCommand("dev", args, config, &globalFlags); err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			return
		}
	},
}

// audit commands
var auditCmd = &cobra.Command{
	Use:     "audit",
	Short:   "Audit the dependencies",
	Long:    "Audit the dependencies of the package manager.",
	GroupID: "dependenciesGroup",
}

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan for vulnerabilities",
	Long:  "Scan for vulnerabilities in the dependencies of the package manager.",
	Args:  cobra.NoArgs,
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
		if err := pm.ExecuteCommand("scan", args, config, &globalFlags); err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			return
		}
	},
}

var fixVulnerabilitiesCmd = &cobra.Command{
	Use:   "fix",
	Short: "Fix the vulnerabilities",
	Long:  "Fix the vulnerabilities of the package manager.",
	Args:  cobra.NoArgs,
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
		if err := pm.ExecuteCommand("fix", args, config, &globalFlags); err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			return
		}
	},
}

// uninstall commands
var removeCmd = &cobra.Command{
	Use:     "remove [dependencies]",
	Short:   "Remove dependencies to the project",
	Long:    "Remove dependencies to the project using the specified package manager.",
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
		if err := pm.ExecuteCommand("remove", args, config, &globalFlags); err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			return
		}
	},
}

// update commands
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

// list commands
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all dependencies",
	Long:    "List all dependencies using the specified package manager.",
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
		if err := pm.ExecuteCommand("list", args, config, &globalFlags); err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			return
		}
	},
}

var outdatedCmd = &cobra.Command{
	Use:     "outdated",
	Short:   "Show outdated dependencies",
	Long:    "Show outdated dependencies using the specified package manager.",
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
		if err := pm.ExecuteCommand("outdated", args, config, &globalFlags); err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			return
		}
	},
}

// cache commands
var cacheManageCmd = &cobra.Command{
	Use:     "cache",
	Short:   "Manage the cache",
	Long:    "Manage the cache using the specified package manager.",
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
		if err := pm.ExecuteCommand("cache", args, config, &globalFlags); err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(localCmd)
	rootCmd.AddCommand(devCmd)
	auditCmd.AddCommand(scanCmd)
	auditCmd.AddCommand(fixVulnerabilitiesCmd)
	rootCmd.AddCommand(auditCmd)
	rootCmd.AddCommand(uninstallCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(upgradeCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(outdatedCmd)
	rootCmd.AddCommand(cacheManageCmd)
}
