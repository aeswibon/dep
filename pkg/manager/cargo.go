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

	commandMap := map[string][]string{
		// basic commands
		"add":     {"add"},    // install a package
		"local":   {"build"},  // install all dependencies
		"remove":  {"remove"}, // remove a package
		"upgrade": {"update"}, // update all packages

		// project management commands
		"init":   {"init"},  // create a new package.json file
		"script": {},        // run a script
		"test":   {"test"},  // run tests
		"build":  {"build"}, // build the project

		// dependency & analysis commands
		"list":     {"list"},              // list all dependencies
		"outdated": {"outdated"},          // list outdated dependencies
		"scan":     {"audit"},             // audit the project for vulnerabilities
		"lock":     {"generate-lockfile"}, // lock dependencies
		"prune":    {"prune"},             // remove extraneous packages

		// configuration & cache commands
		"config": {"config"}, // view or set configuration
		"cache":  {"cache"},  // manage the cache

		// workspace commands
		"workspace-ls": {"metadata", "--format-version=1"}, // list workspaces

		// publish commands
		"publish": {"publish"}, // publish a package
		"login":   {"login"},   // login to the registry
		"pack":    {"package"}, // create a distributable package
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
