package cmd

import "github.com/spf13/cobra"

var dependenciesGroup = &cobra.Group{
	ID:    "dependenciesGroup",
	Title: "Manage Dependencies",
}

var scriptGroup = &cobra.Group{
	ID:    "scriptGroup",
	Title: "Run Scripts",
}
