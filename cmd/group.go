package cmd

import "github.com/spf13/cobra"

var dependenciesGroup = &cobra.Group{
	ID:    "dependenciesGroup",
	Title: "Manage Dependencies",
}