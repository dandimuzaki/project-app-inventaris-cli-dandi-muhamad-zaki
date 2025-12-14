package cmd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/handler"
	"github.com/spf13/cobra"
)

// Initialize input fields
var (
	categoryNameStr string
	categoryName    sql.NullString
	descriptionStr string
	description sql.NullString
	categoryIdInt int

	itemIdInt int
	itemNameStr string
	itemName sql.NullString
	categoryId sql.NullInt32
	priceFl float64
	price sql.NullFloat64
	dateStr string
	purchaseDate sql.NullTime

	keyword string

	rootCmd = &cobra.Command{
		Use: "inv",
		Short: "Simple Inventory Management with Golang CLI",
	}
)

func Execute() {
	// Execute root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Init(handler handler.Handler) {
	// Initialize necessary commands
	initCategoryCmd(handler)
	initItemCmd(handler)
	initReportCmd(handler)
}