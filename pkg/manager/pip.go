package manager

import (
	"fmt"

	"github.com/aeswibon/dep/pkg/types"
	"github.com/aeswibon/dep/pkg/utils"
)

// PipPackageManager represents the Go package manager
type PipPackageManager struct{}

// ExecuteCommand executes the given command with the provided arguments
func (p *PipPackageManager) ExecuteCommand(command string, args []string, config *types.Config, flags *types.GlobalFlags) error {
	var cmdArgs []string

	commandMap := map[string][]string{
		// basic commands
		"add":       {"install"},           // install a package
		"local":     {"install"},           // install all dependencies
		"install":   {"install", "--user"}, // install a package globally
		"dev":       {"install", "--dev"},  // install a package as a dev dependency
		"remove":    {"uninstall"},         // remove a package
		"uninstall": {"uninstall"},         // remove a package globally
		"update":    {"install", "-U"},     // update a package
		"upgrade":   {"install", "-U"},     // update all packages

		// project management commands
		"init":   {"init"},  // create a new setup.py file
		"script": {},        // run a script
		"test":   {"test"},  // run tests
		"build":  {"build"}, // build the project

		// dependency & analysis commands
		"list":     {"list"},               // list all dependencies
		"outdated": {"list", "--outdated"}, // list outdated dependencies
		"scan":     {"check"},              // audit the project for vulnerabilities
		"lock":     {"freeze"},             // lock dependencies
		"detect":   {"show", "-f"},         // detect why a package is needed
		"prune":    {"uninstall", "-y"},    // remove extraneous packages

		// configuration & cache commands
		"config": {"config"}, // view or set configuration
		"cache":  {"cache"},  // manage the cache

		// workspace commands
		"workspace": {"workspace"}, // manage workspaces

		// publish commands
		"publish": {"setup.py", "upload"},   // publish a package
		"login":   {"setup.py", "register"}, // login to the registry
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
func (p *PipPackageManager) GetCommand() string {
	return "pip"
}
