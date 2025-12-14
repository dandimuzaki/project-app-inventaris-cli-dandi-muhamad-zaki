package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/dto"
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/handler"
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/utils"
	"github.com/spf13/cobra"
)

func initItemCmd(handler handler.Handler) {
	var listItem = &cobra.Command{
		Use: "ls-item",
		Short: "View list of all items",
		Run: func(cmd *cobra.Command, args []string) {
			// Retrieve items data without keyword and categoryId
			d := dto.ItemParam{keyword, categoryId}
			data, err := handler.HandlerItem.GetItems(d)
			if err != nil {
				utils.PrintErr(err)
			} else {
				utils.PrintSuccess("Item list is retrieved")
				utils.TableItem(data)
			}
		},
	}
	rootCmd.AddCommand(listItem)

	var replacedItem = &cobra.Command{
		Use: "replace-item",
		Short: "View list of all items that must be replaced after 100 days of usage",
		Run: func(cmd *cobra.Command, args []string) {
			// Retrieve items that must be replaced
			data, err := handler.HandlerItem.MustReplacedItems()
			if err != nil {
				utils.PrintErr(err)
			} else {
				utils.PrintSuccess("List of items must be replaced is retrieved")
				utils.TableItem(data)
			}
		},
	}
	rootCmd.AddCommand(replacedItem)

	var searchItem = &cobra.Command{
		Use: "sc-item",
		Short: "Search item based on name",
		Run: func(cmd *cobra.Command, args []string) {
			// Retrieve items data with keyword based on item name
			d := dto.ItemParam{keyword, categoryId}
			data, err := handler.HandlerItem.GetItems(d)
			if err != nil {
				utils.PrintErr(err)
			} else {
				fmt.Printf("Showing %v results for %s :\n", len(data), keyword)
				utils.TableItem(data)
			}
		},
	}
	searchItem.Flags().StringVarP(&keyword, "keyword", "k", "", "keyword for searching item based on name")
	rootCmd.AddCommand(searchItem)

	var filterItem = &cobra.Command{
		Use: "flt-item",
		Short: "Filter item based on category",
		Run: func(cmd *cobra.Command, args []string) {
			// Check if category id is valid
			categories, msg := handler.HandlerCategory.GetCategories()
			if msg != nil {
				utils.PrintErr(msg)
				return
			}
			found := false
			for _, c := range categories {
				if c.ID == categoryIdInt {
					found = true
				}
			}
			if !found {
				utils.PrintErr(errors.New("category id is not valid"))
				return
			}

			// Make categoryId not null
			if categoryIdInt != 0 {
				categoryId.Int32 = int32(categoryIdInt)
				categoryId.Valid = true
			}

			// Retrieve items data with category filter based on category id
			d := dto.ItemParam{keyword, categoryId}
			data, err := handler.HandlerItem.GetItems(d)
			if err != nil {
				utils.PrintErr(err)
			} else if len(data) > 0 {
				fmt.Printf("Showing %v items for category %s :\n", len(data), data[0].Category)
				utils.TableItem(data)
			}
		},
	}
	filterItem.Flags().IntVarP(&categoryIdInt, "categoryId", "c", 0, "category id")
	rootCmd.AddCommand(filterItem)

	var viewItem = &cobra.Command{
		Use: "v-item",
		Short: "View an item by ID",
		Run: func(cmd *cobra.Command, args []string) {
			// Item id should be greater than 0
			if itemIdInt<=0 {
				utils.PrintErr(errors.New("id must be greater than 0 and is required"))
				return
			}

			// Process retrieving item data based on item id
			data, err := handler.HandlerItem.GetItemByID(itemIdInt)
			if err != nil {
				utils.PrintErr(err)
			} else {
				utils.CardItem(*data)
			}
		},
	}
	viewItem.Flags().IntVarP(&itemIdInt, "id", "i", 0, "item id")
	rootCmd.AddCommand(viewItem)

	var createItem = &cobra.Command{
		Use: "cre-item",
		Short: "Create a new item",
		Run: func(cmd *cobra.Command, args []string) {
			// Check for empty input
			if itemNameStr == "" {
				utils.PrintErr(errors.New("item name is required"))
				return
			}

			if categoryIdInt == 0 {
				utils.PrintErr(errors.New("category id is required"))
				return
			}

			if priceFl == 0 {
				utils.PrintErr(errors.New("price is required"))
				return
			}

			if dateStr == "" {
				utils.PrintErr(errors.New("purchase date is required"))
				return
			}

			// Check if category id is valid
			categories, msg := handler.HandlerCategory.GetCategories()
			if msg != nil {
				utils.PrintErr(msg)
				return
			}
			found := false
			for _, c := range categories {
				if c.ID == categoryIdInt {
					found = true
				}
			}
			if !found {
				utils.PrintErr(errors.New("category id is not valid"))
				return
			}

			// Parse time input into datetime
			date, err := time.Parse("2006-01-02", dateStr)
			if err != nil {
				utils.PrintErr(err)
				return
			}

			// Process the create item request
			req := dto.CreateItemRequest{itemNameStr, categoryIdInt, priceFl, date}
			data, err := handler.HandlerItem.CreateItem(req)
			if err != nil {
				utils.PrintErr(err)
			} else {
				utils.PrintSuccess("New item is created")
				utils.CardItem(*data)
			}
		},
	}

	createItem.Flags().StringVarP(&itemNameStr, "name", "n", "", "item name")
	createItem.Flags().IntVarP(&categoryIdInt, "categoryId", "c", 0, "category id")
	createItem.Flags().Float64VarP(&priceFl, "price", "p", 0, "item price")
	createItem.Flags().StringVarP(&dateStr, "purchaseDate", "d", "", "item purchase date (YYYY-MM-DD)")
	rootCmd.AddCommand(createItem)

	var updateItem = &cobra.Command{
		Use: "upd-item",
		Short: "Update name or description of an item",
		Run: func(cmd *cobra.Command, args []string) {
			// Item id should be greater than 0
			if itemIdInt<=0 {
				utils.PrintErr(errors.New("id must be greater than 0 and is required"))
				return
			}

			// Check if category id is valid
			categories, msg := handler.HandlerCategory.GetCategories()
			if msg != nil {
				utils.PrintErr(msg)
				return
			}
			found := false
			for _, c := range categories {
				if c.ID == categoryIdInt {
					found = true
				}
			}
			if categoryIdInt != 0 && !found {
				utils.PrintErr(errors.New("category id is not valid"))
				return
			}

			// Item is not updated if no field is entered
			if itemNameStr == "" && categoryIdInt == 0 && priceFl == 0 && dateStr == "" {
				utils.PrintErr(errors.New("no field is updated"))
				return
			}

			// Make itemName not null
			if itemNameStr != "" {
				itemName.String = itemNameStr
				itemName.Valid = true
			}

			// Make categoryId not null
			if categoryIdInt != 0 {
				categoryId.Int32 = int32(categoryIdInt)
				categoryId.Valid = true
			}

			// Make price not null
			if priceFl != 0 {
				price.Float64 = priceFl
				price.Valid = true
			}

			// Make purchaseDate not null
			if dateStr != "" {
				// Parse the purchase date string into datetime
				date, err := time.Parse("2006-01-02", dateStr)
				if err != nil {
					utils.PrintErr(err)
					return
				}
				purchaseDate.Time = date
				purchaseDate.Valid = true
			}

			// Process the update item request
			d := dto.UpdateItemRequest{itemName, categoryId, price, purchaseDate}
			err := handler.HandlerItem.UpdateItem(itemIdInt, d)
			if err != nil {
				utils.PrintErr(msg)
			} else {
				utils.PrintSuccess(fmt.Sprintf("Item %d is updated", itemIdInt))
			}
		},
	}
	updateItem.Flags().IntVarP(&itemIdInt, "id", "i", 0, "item id")
	updateItem.Flags().StringVarP(&itemNameStr, "name", "n", "", "item name")
	updateItem.Flags().IntVarP(&categoryIdInt, "categoryId", "c", 0, "category id")
	updateItem.Flags().Float64VarP(&priceFl, "price", "p", 0, "item price")
	updateItem.Flags().StringVarP(&dateStr, "purchaseDate", "d", "", "item purchase date (YYYY-MM-DD)")
	rootCmd.AddCommand(updateItem)

	var deleteItem = &cobra.Command{
		Use: "del-item",
		Short: "Delete an item by ID",
		Run: func(cmd *cobra.Command, args []string) {
			// Item id should be greater than 0
			if itemIdInt<=0 {
				utils.PrintErr(errors.New("id must be greater than 0 and is required"))
				return
			}

			// Process the delete item request
			err := handler.HandlerItem.DeleteItem(itemIdInt)
			if err != nil {
				utils.PrintErr(err)
			} else {
				utils.PrintSuccess(fmt.Sprintf("Item %d is deleted", itemIdInt))
			}
		},
	}
	deleteItem.Flags().IntVarP(&itemIdInt, "id", "i", 0, "item id")
	rootCmd.AddCommand(deleteItem)
}