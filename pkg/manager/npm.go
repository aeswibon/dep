package manager

import (
	"fmt"

	"github.com/aeswibon/dep/pkg/types"
	"github.com/aeswibon/dep/pkg/utils"
)

// NpmPackageManager represents the NPM package manager
type NpmPackageManager struct{}

// ExecuteCommand executes the given command with the provided arguments
func (n *NpmPackageManager) ExecuteCommand(command string, args []string, config *types.Config, flags *types.GlobalFlags) error {
	var cmdArgs []string
	// switch statement to handle different commands
	commandMap := map[string][]string{
		"install":         {"install"},
		"add":             {"install"},
		"dev":             {"install", "--save-dev"},
		"global":          {"install", "-g"},
		"remove":          {"uninstall"},
		"remove-global":   {"uninstall", "-g"},
		"update":          {"update"},
		"upgrade":         {"update"},
		"init":            {"init", "-y"},
		"run-script":      {"run"},
		"tests":           {"test"},
		"list":            {"list"},
		"outdated":        {"outdated"},
		"cache-clean":     {"cache", "clean"},
		"cache-verify":    {"cache", "verify"},
		"publish-package": {"publish"},
		"login":           {"login"},
		"audit":           {"audit"},
		"fix-audit":       {"audit", "fix"},
	}

	// Check if the command exists in the map
	if cmd, exists := commandMap[command]; exists {
		cmdArgs = append(cmd, args...)
	} else if cmd := utils.CheckCommand(command, config); cmd != "" {
		cmdArgs = append([]string{cmd}, args...)
	} else {
		fmt.Printf("Unsupported command: %s\n", command)
		return nil
	}

	return utils.ExecuteCommand(n.GetCommand(), cmdArgs, config, flags)
}

// GetCommand returns the command for the package manager
func (n *NpmPackageManager) GetCommand() string {
	return "npm"
}
