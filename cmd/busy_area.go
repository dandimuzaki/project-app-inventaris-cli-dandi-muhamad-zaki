package cmd

import (
	"fmt"
	"os"
	"session-14/handler"
)

func BusyAreas(handlerReport handler.HandlerReport) {
	var choice string
	handlerReport.BusyAreas()
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
