package cmd

import (
	"fmt"

	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/handler"
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/utils"
	"github.com/spf13/cobra"
)

func initReportCmd(handler handler.Handler) {
	var totalNetValue = &cobra.Command{
		Use: "value",
		Short: "Total net value of items after depreciation",
		Run: func(cmd *cobra.Command, args []string) {
			// Retrieve total net value of items after calculation of depreciation
			data, err := handler.HandlerReport.TotalNetValue()
			if err != nil {
				utils.PrintErr(err)
			} else {
				fmt.Printf("\033[33mTotal net value of your inventory is Rp %.2f\033[0m", *data)
			}
		},
	}
	rootCmd.AddCommand(totalNetValue)

	var totalInvestment = &cobra.Command{
		Use: "invest",
		Short: "Total price of items invested",
		Run: func(cmd *cobra.Command, args []string) {
			// Retrieve total price of items invested in inventory
			data, err := handler.HandlerReport.TotalInvestment()
			if err != nil {
				utils.PrintErr(err)
			} else {
				fmt.Printf("\033[33mTotal investment of your inventory is Rp %.2f\033[0m", *data)
			}
		},
	}
	rootCmd.AddCommand(totalInvestment)
}