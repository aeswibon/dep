package utils

import (
	"fmt"
	"os/exec"

	"github.com/aeswibon/dep/pkg/types"
)

// CheckTool checks if the given command is available in the system.
func CheckTool(command string) error {
	_, err := exec.LookPath(command)
	if err != nil {
		return fmt.Errorf("command not found: %s", command)
	}
	return nil
}

// CheckCommand checks if the given command is present in the config file
func CheckCommand(command string, config *types.Config) string {
	for _, cmd := range config.CustomCommands {
		if cmd.Name == command {
			return cmd.Command
		}
	}
	return ""
}
