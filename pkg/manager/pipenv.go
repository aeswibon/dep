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

	commandMap := map[string][]string{
		// basic commands
		"add":     {"install"},          // install a package
		"local":   {"install"},          // install all dependencies
		"dev":     {"install", "--dev"}, // install a package as a dev dependency
		"remove":  {"uninstall"},        // remove a package
		"update":  {"update"},           // update a package
		"upgrade": {"update"},           // update all packages

		// project management commands
		"init":   {},               // create a new Pipfile
		"script": {"run"},          // run a script
		"test":   {"run", "test"},  // run tests
		"build":  {"run", "build"}, // build the project

		// dependency & analysis commands
		"list":     {"graph"},                // list all dependencies
		"outdated": {"update", "--outdated"}, // list outdated dependencies
		"scan":     {"check"},                // audit the project for vulnerabilities
		"fix":      {"check", "--outdated"},  // fix vulnerabilities
		"dedupe":   {"install"},              // remove duplicate dependencies
		"lock":     {"lock"},                 // lock dependencies
		"prune":    {"clean"},                // remove extraneous packages

		// configuration & cache commands
		"config": {"check"}, // view or set configuration
		"cache":  {"cache"}, // manage the cache
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
