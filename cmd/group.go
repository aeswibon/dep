package cmd

import "github.com/spf13/cobra"

var globalGroup = &cobra.Group{
	ID:    "globalGroup",
	Title: "Global Commands",
}

var dependenciesGroup = &cobra.Group{
	ID:    "dependenciesGroup",
	Title: "Manage Project Dependencies",
}

var scriptGroup = &cobra.Group{
	ID:    "scriptGroup",
	Title: "Run Scripts",
}
