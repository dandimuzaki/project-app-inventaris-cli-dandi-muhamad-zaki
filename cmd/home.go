package cmd

import (
	"fmt"
	"os"
)

func HomePage() {
	for {
		fmt.Println("\n=== HOME MENU ===")
		fmt.Println("1. Create Data")
		fmt.Println("2. List Data")
		fmt.Println("3. Update Data")
		fmt.Println("4. Delete Data")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:

		case 2:

		case 3:

		case 4:

		case 5:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please choose between 1 - 5.")
		}
	}

}
