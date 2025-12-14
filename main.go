package main

import (
	"context"
	"log"

	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/cmd"
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/database"
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/handler"
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/repository"
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/service"
)

func main() {
	// Init DB connection
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())
	repo := repository.NewRepository(db) 		// Init repo
	service := service.NewService(repo) 		// Init service
	handler := handler.NewHandler(service) 	// Init handler
	cmd.Init(handler)												// Init command
	cmd.Execute()														// Execute command
}
