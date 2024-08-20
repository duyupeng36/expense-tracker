package cmd

import (
	"expense-tracker/app"
	"github.com/spf13/cobra"
	"log"
)

var description string
var amount float64

func init() {
	addCmd.Flags().StringVarP(&description, "description", "", "", "expense description")
	err := addCmd.MarkFlagRequired("description")
	if err != nil {
		log.Fatal(err)
	}
	addCmd.Flags().Float64VarP(&amount, "amount", "", 0.0, "expense amount")
	err = addCmd.MarkFlagRequired("amount")
	if err != nil {
		log.Fatal(err)
	}
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add expense",
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.App.Add(description, amount)
	},
}
