package cmd

import (
	"expense-tracker/app"
	"github.com/spf13/cobra"
	"time"
)

var month int

func init() {
	summaryCmd.Flags().IntVarP(&month, "month", "m", 0, "summary month")
}

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "summary expense of mount or total year",
	Run: func(cmd *cobra.Command, args []string) {
		app.App.Summary(time.Month(month))
	},
}
