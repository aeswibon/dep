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

	commandMap := map[string][]string{
		// basic commands
		"add":     {"get"},                // install a package
		"local":   {"mod", "download"},    // install all dependencies
		"install": {"install"},            // install a package globally
		"remove":  {"mod", "tidy"},        // remove a package
		"update":  {"get", "-u"},          // update a package
		"upgrade": {"get", "-u", "./..."}, // update all packages

		// project management commands
		"init":   {"mod", "init"}, // create a new go.mod file
		"script": {"run"},         // run a script
		"test":   {"test"},        // run tests
		"build":  {"build"},       // build the project

		// dependency & analysis commands
		"list":     {"list", "-m", "all"},       // list all dependencies
		"outdated": {"list", "-u", "-m", "all"}, // list outdated dependencies
		"scan":     {"vet"},                     // audit the project for vulnerabilities
		"lock":     {"mod", "vendor"},           // lock dependencies
		"detect":   {"mod", "verify"},           // detect why a package is needed
		"prune":    {"mod", "tidy"},             // remove extraneous packages

		// configuration & cache commands
		"config": {"env"},   // view or set configuration
		"cache":  {"clean"}, // manage the cache
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
