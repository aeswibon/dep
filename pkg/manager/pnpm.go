package manager

import (
	"fmt"

	"github.com/aeswibon/dep/pkg/types"
	"github.com/aeswibon/dep/pkg/utils"
)

// PnpmPackageManager represents the NPM package manager
type PnpmPackageManager struct{}

// ExecuteCommand executes the given command with the provided arguments
func (p *PnpmPackageManager) ExecuteCommand(command string, args []string, config *types.Config, flags *types.GlobalFlags) error {
	var cmdArgs []string
	// switch statement to handle different commands
	commandMap := map[string][]string{
		"install":         {"install"},
		"add":             {"add"},
		"dev":             {"add", "-D"},
		"global":          {"add", "-g"},
		"remove":          {"remove"},
		"remove-global":   {"global", "remove"},
		"update":          {"update"},
		"upgrade":         {"upgrade"},
		"init":            {"init", "-y"},
		"run-script":      {},
		"tests":           {"test"},
		"list":            {"list"},
		"outdated":        {"outdated"},
		"cache-clean":     {"store", "prune"},
		"cache-verify":    {"store", "verify"},
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

	return utils.ExecuteCommand(p.GetCommand(), cmdArgs, config, flags)
}

// GetCommand returns the command for the package manager
func (p *PnpmPackageManager) GetCommand() string {
	return "pnpm"
}
