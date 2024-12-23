package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aeswibon/dep/pkg/types"
)

// LoadConfig loads the configuration file from the current directory
func LoadConfig() (*types.Config, error) {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get current directory: %v", err)
	}

	// Get the path to the config file
	filepath := fmt.Sprintf("%s/config.json", currentDir)

	// Open the config file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}

	defer file.Close()

	// Decode the config file
	var c types.Config
	if err := json.NewDecoder(file).Decode(&c); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %v", err)
	}

	return &c, nil
}
