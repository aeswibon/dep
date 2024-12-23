package manager

import (
	"fmt"

	"github.com/aeswibon/dep/pkg/types"
	"github.com/aeswibon/dep/pkg/utils"
)

// GoPackageManager represents the Go package manager
type GoPackageManager struct{}

// ExecuteCommand executes the given command with the provided arguments
func (g *GoPackageManager) ExecuteCommand(command string, args []string, config *types.Config, flags *types.GlobalFlags) error {
	var cmdArgs []string
	// switch statement to handle different commands
	commandMap := map[string][]string{
		"install":     {"mod", "download"},
		"add":         {"get"},
		"global":      {"install"},
		"remove":      {"mod", "tidy"},
		"update":      {"get", "-u"},
		"upgrade":     {"get", "-u", "./..."},
		"init":        {"mod", "init"},
		"run-script":  {"run"},
		"tests":       {"test"},
		"list":        {"list", "-m", "all"},
		"outdated":    {"list", "-u", "-m", "all"},
		"cache-clean": {"clean", "-cache"},
		"audit":       {"vet"},
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

	return utils.ExecuteCommand(g.GetCommand(), cmdArgs, config, flags)
}

// GetCommand returns the command for the package manager
func (g *GoPackageManager) GetCommand() string {
	return "go"
}
