package cmd

import (
	"expense-tracker/app"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all expense",
	Run: func(cmd *cobra.Command, args []string) {
		app.App.List()
	},
}
