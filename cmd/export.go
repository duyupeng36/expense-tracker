package cmd

import (
	"expense-tracker/fs"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "export the expense to csv file",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fs.Export2CSV()
	},
}
