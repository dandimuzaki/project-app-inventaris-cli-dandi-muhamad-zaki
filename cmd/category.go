package cmd

import (
	"errors"
	"fmt"

	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/dto"
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/handler"
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/utils"
	"github.com/spf13/cobra"
)

func initCategoryCmd(handler handler.Handler) {
	var listCategory = &cobra.Command{
		Use: "ls-cat",
		Short: "View list of categories",
		Run: func(cmd *cobra.Command, args []string) {
			data, err := handler.HandlerCategory.GetCategories()
			if err != nil {
				utils.PrintErr(err)
			} else {
				utils.PrintSuccess("Category list is retrieved")
				utils.TableCategory(data)
			}
		},
	}
	rootCmd.AddCommand(listCategory)

	var viewCategory = &cobra.Command{
		Use: "v-cat",
		Short: "View a category by ID",
		Run: func(cmd *cobra.Command, args []string) {
			// Category id should be greater than 0
			if categoryIdInt<=0 {
				utils.PrintErr(errors.New("id must be greater than 0 and is required"))
				return
			}

			// Process retrieving category data based on category id
			data, err := handler.HandlerCategory.GetCategoryByID(categoryIdInt)
			if err != nil {
				utils.PrintErr(err)
			} else {
				utils.CardCategory(*data)
			}
		},
	}
	viewCategory.Flags().IntVarP(&categoryIdInt, "id", "i", 0, "category id")
	rootCmd.AddCommand(viewCategory)

	var createCategory = &cobra.Command{
		Use: "cre-cat",
		Short: "Create a new category",
		Run: func(cmd *cobra.Command, args []string) {
			// Check for empty category name
			if categoryNameStr == "" {
				utils.PrintErr(errors.New("category name is required"))
				return
			}

			// Make description not null if user enter description
			if descriptionStr != "" {
				description.String = descriptionStr
				description.Valid = true
			}

			// Process the create category request
			req := dto.CreateCategoryRequest{categoryNameStr, description}
			data, err := handler.HandlerCategory.CreateCategory(req)
			if err != nil {
				utils.PrintErr(err)
			} else {
				utils.PrintSuccess("New category is created")
				utils.CardCategory(*data)
			}
		},
	}

	createCategory.Flags().StringVarP(&categoryNameStr, "name", "n", "", "category name")
	createCategory.Flags().StringVarP(&descriptionStr, "description", "d", "", "category description")
	rootCmd.AddCommand(createCategory)

	var updateCategory = &cobra.Command{
		Use: "upd-cat",
		Short: "Update name or description of a category",
		Run: func(cmd *cobra.Command, args []string) {
			// Category id should be greater than 0
			if categoryIdInt<=0 {
				utils.PrintErr(errors.New("id must be greater than 0 and is required"))
				return
			}

			// Category is not updated if no field is entered
			if categoryNameStr == "" && descriptionStr == "" {
				utils.PrintErr(errors.New("no field is updated"))
				return
			}

			// Make categoryName not null
			if categoryNameStr != "" {
				categoryName.String = categoryNameStr
				categoryName.Valid = true
			}

			// Make description not null
			if descriptionStr != "" {
				description.String = descriptionStr
				description.Valid = true
			}

			// Process the update category request
			d := dto.UpdateCategoryRequest{categoryName, description}
			err := handler.HandlerCategory.UpdateCategory(categoryIdInt, d)
			if err != nil {
				utils.PrintErr(err)
			} else {
				utils.PrintSuccess(fmt.Sprintf("Category %d is updated", categoryIdInt))
			}
		},
	}
	updateCategory.Flags().IntVarP(&categoryIdInt, "id", "i", 0, "category id")
	updateCategory.Flags().StringVarP(&categoryNameStr, "name", "n", "", "category name")
	updateCategory.Flags().StringVarP(&descriptionStr, "desciption", "d", "", "category description")
	rootCmd.AddCommand(updateCategory)

	var deleteCategory = &cobra.Command{
		Use: "del-cat",
		Short: "Delete a category by ID",
		Run: func(cmd *cobra.Command, args []string) {
			// Category id should be greater than 0
			if categoryIdInt<=0 {
				utils.PrintErr(errors.New("id must be greater than 0 and is required"))
				return
			}

			// Process the delete category request
			err := handler.HandlerCategory.DeleteCategory(categoryIdInt)
			if err != nil {
				utils.PrintErr(err)
			} else {
				utils.PrintSuccess(fmt.Sprintf("Category %d is deleted", categoryIdInt))
			}
		},
	}
	deleteCategory.Flags().IntVarP(&categoryIdInt, "id", "i", 0, "category id")
	rootCmd.AddCommand(deleteCategory)
}