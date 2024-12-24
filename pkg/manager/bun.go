package manager

import (
	"fmt"

	"github.com/aeswibon/dep/pkg/types"
	"github.com/aeswibon/dep/pkg/utils"
)

// BunPackageManager represents the NPM package manager
type BunPackageManager struct{}

// ExecuteCommand executes the given command with the provided arguments
func (b *BunPackageManager) ExecuteCommand(command string, args []string, config *types.Config, flags *types.GlobalFlags) error {
	var cmdArgs []string
	// switch statement to handle different commands
	commandMap := map[string][]string{
		"add":           {"add"},
		"cache-clean":   {"pm", "cache", "clear"},
		"dev":           {"add", "-D"},
		"global":        {"add", "-g"},
		"init":          {"init", "-y"},
		"install":       {"install"},
		"list":          {"pm", "ls"},
		"outdated":      {"pm", "outdated"},
		"package-info":  {"pm", "info"},
		"rebuild-deps":  {"rebuild"},
		"remove":        {"remove"},
		"remove-global": {"global", "remove"},
		"run-script":    {"run"},
		"tests":         {"test"},
		"update":        {"update"},
		"upgrade":       {"upgrade"},
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

	return utils.ExecuteCommand(b.GetCommand(), cmdArgs, config, flags)
}

// GetCommand returns the command for the package manager
func (b *BunPackageManager) GetCommand() string {
	return "bun"
}
