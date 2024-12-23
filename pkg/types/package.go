package types

// PackageManager interface defines the methods that a package manager must implement
type PackageManager interface {
	ExecuteCommand(command string, args []string, config *Config, flags *GlobalFlags) error
	GetCommand() string
}
