package manager

import (
	"fmt"

	"github.com/aeswibon/dep/pkg/types"
	"github.com/aeswibon/dep/pkg/utils"
)

// CargoPackageManager represents the Go package manager
type CargoPackageManager struct{}

// ExecuteCommand executes the given command with the provided arguments
func (c *CargoPackageManager) ExecuteCommand(command string, args []string, config *types.Config, flags *types.GlobalFlags) error {
	var cmdArgs []string
	// switch statement to handle different commands
	commandMap := map[string][]string{
		"install":         {"build"},
		"add":             {"add"},
		"dev":             {"add", "--dev"},
		"global":          {"install"},
		"remove":          {"remove"},
		"update":          {"update"},
		"upgrade":         {"update"},
		"init":            {"init"},
		"run-script":      {"run"},
		"test":            {"test"},
		"list":            {"tree"},
		"outdated":        {"outdated"},
		"cache-clean":     {"clean"},
		"publish-package": {"publish"},
		"login":           {"login"},
		"audit":           {"audit"},
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

	return utils.ExecuteCommand(c.GetCommand(), cmdArgs, config, flags)
}

// GetCommand returns the command for the package manager
func (c *CargoPackageManager) GetCommand() string {
	return "cargo"
}
