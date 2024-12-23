package manager

import (
	"fmt"

	"github.com/aeswibon/dep/pkg/types"
	"github.com/aeswibon/dep/pkg/utils"
)

// PipenvPackageManager represents the Go package manager
type PipenvPackageManager struct{}

// ExecuteCommand executes the given command with the provided arguments
func (p *PipenvPackageManager) ExecuteCommand(command string, args []string, config *types.Config, flags *types.GlobalFlags) error {
	var cmdArgs []string
	// switch statement to handle different commands
	commandMap := map[string][]string{
		"install":     {"install"},
		"add":         {"install"},
		"dev":         {"install", "--dev"},
		"remove":      {"uninstall"},
		"update":      {"update"},
		"upgrade":     {"update"},
		"init":        {},
		"run-script":  {"run"},
		"test":        {"run", "test"},
		"list":        {"graph"},
		"outdated":    {"update", "--outdated"},
		"cache-clean": {"clean"},
		"audit":       {"check"},
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
func (p *PipenvPackageManager) GetCommand() string {
	return "pipenv"
}
