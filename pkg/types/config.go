package types

type CustomCommand struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Command     string `json:"command"`
}

// Config represents the structure of the config.json file
type Config struct {
	PackageManager string          `json:"packageManager"`
	CustomCommands []CustomCommand `json:"customCommands"`
}
