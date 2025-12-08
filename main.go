package main

import (
	"context"
	"log"
	"session-14/cmd"
	"session-14/database"
	"session-14/handler"
	"session-14/repository"
	"session-14/service"
)

func main() {
	// init DB connection
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())

	// init object
	repoReport := repository.NewRepoReport(db)
	serviceReport := service.NewServiceReport(&repoReport)
	hadlerReport := handler.NewHandlerReport(&serviceReport)
	cmd.HomePage(hadlerReport)
}
