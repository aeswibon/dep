package manager

import (
	"fmt"

	"github.com/aeswibon/dep/pkg/types"
	"github.com/aeswibon/dep/pkg/utils"
)

// NpmPackageManager represents the NPM package manager
type NpmPackageManager struct{}

// ExecuteCommand executes the given command with the provided arguments
func (n *NpmPackageManager) ExecuteCommand(command string, args []string, config *types.Config, flags *types.GlobalFlags) error {
	var cmdArgs []string

	commandMap := map[string][]string{
		// basic commands
		"add":       {"install"},               // install a package
		"local":     {"install"},               // install all dependencies
		"install":   {"install", "-g"},         // install a package globally
		"dev":       {"install", "--save-dev"}, // install a package as a dev dependency
		"remove":    {"uninstall"},             // remove a package
		"uninstall": {"uninstall", "-g"},       // remove a package globally
		"update":    {"update"},                // update a package
		"upgrade":   {"update"},                // update all packages

		// project management commands
		"init":    {"init", "-y"}, // create a new package.json file
		"script":  {"run"},        // run a script
		"test":    {"test"},       // run tests
		"build":   {"build"},      // build the project
		"version": {"version"},    // bump the version

		// dependency & analysis commands
		"list":     {"list"},         // list all dependencies
		"outdated": {"outdated"},     // list outdated dependencies
		"scan":     {"audit"},        // audit the project for vulnerabilities
		"fix":      {"audit", "fix"}, // fix vulnerabilities
		"dedupe":   {"dedupe"},       // remove duplicate dependencies
		"lock":     {"shrinkwrap"},   // lock dependencies
		"prune":    {"prune"},        // remove extraneous packages

		// configuration & cache commands
		"config": {"config"}, // view or set configuration
		"cache":  {"cache"},  // manage the cache

		// workspace commands
		"workspace-ls": {"query", ".workspaces"}, // list workspaces

		// publish commands
		"publish":   {"publish"},   // publish a package
		"login":     {"login"},     // login to the registry
		"pack":      {"pack"},      // create a tarball
		"deprecate": {"deprecate"}, // deprecate a package
		"unpublish": {"unpublish"}, // unpublish a package
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

	return utils.ExecuteCommand(n.GetCommand(), cmdArgs, config, flags)
}

// GetCommand returns the command for the package manager
func (n *NpmPackageManager) GetCommand() string {
	return "npm"
}
