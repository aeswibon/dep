package manager

import (
	"fmt"

	"github.com/aeswibon/dep/pkg/types"
	"github.com/aeswibon/dep/pkg/utils"
)

// YarnPackageManager represents the NPM package manager
type YarnPackageManager struct{}

// ExecuteCommand executes the given command with the provided arguments
func (y *YarnPackageManager) ExecuteCommand(command string, args []string, config *types.Config, flags *types.GlobalFlags) error {
	var cmdArgs []string
	// switch statement to handle different commands
	commandMap := map[string][]string{
		"install":         {"install"},
		"add":             {"add"},
		"dev":             {"add", "--dev"},
		"global":          {"add", "global"},
		"remove":          {"remove"},
		"remove-global":   {"global", "remove"},
		"update":          {"upgrade"},
		"upgrade":         {"upgrade"},
		"init":            {"init", "-y"},
		"run-script":      {},
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

	return utils.ExecuteCommand(y.GetCommand(), cmdArgs, config, flags)
}

// GetCommand returns the command for the package manager
func (y *YarnPackageManager) GetCommand() string {
	return "yarn"
}
