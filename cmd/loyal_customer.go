package cmd

import (
	"fmt"
	"os"
	"session-14/handler"
)

func MonthlyLoyalCustomers(handlerReport handler.HandlerReport) {
	var status, choice string
	fmt.Print("Status : ")
	fmt.Scan(&status)
	handlerReport.MonthlyLoyalCustomers(status)
	fmt.Print("Apakah kamu ingin melanjutkan ke halaman lain ? ya / tidak")
	fmt.Scan(&choice)

	switch choice {
	case "ya":
		ClearScreen()
	case "tidak":
		os.Exit(0)
	default:
		fmt.Println("Pilihan kamu salah tolong masukkan ulang")
	}
}
