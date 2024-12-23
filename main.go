package main

import (
	"log"
	"os"

	"github.com/aeswibon/dep/cmd"
)

func main() {
	// Check if the config file exists
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		log.Fatalf("config.json not found in the current directory\n")
	}

	// Execute the root command
	cmd.Execute()
}
