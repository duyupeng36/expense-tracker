package cmd

import (
	"expense-tracker/app"
	"github.com/spf13/cobra"
	"log"
)

var id int

func init() {
	deleteCmd.Flags().IntVarP(&id, "id", "i", 0, "id of the user")
	if err := deleteCmd.MarkFlagRequired("id"); err != nil {
		log.Fatal(err)
	}
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a expense by id",
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.App.Delete(id)
	},
}
