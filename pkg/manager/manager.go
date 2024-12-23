package manager

import (
	"fmt"

	"github.com/aeswibon/dep/pkg/types"
)

// GetPackageManager returns the package manager based on the configuration
func GetPackageManager(manager string) (types.PackageManager, error) {
	switch manager {
	case "npm":
		return &NpmPackageManager{}, nil
	case "go":
		return &GoPackageManager{}, nil
	case "yarn":
		return &YarnPackageManager{}, nil
	case "bun":
		return &BunPackageManager{}, nil
	case "pnpm":
		return &PnpmPackageManager{}, nil
	case "pip":
		return &PipPackageManager{}, nil
	case "pipenv":
		return &PipenvPackageManager{}, nil
	case "cargo":
		return &CargoPackageManager{}, nil
	default:
		return nil, fmt.Errorf("unsupported package manager: %s", manager)
	}
}
