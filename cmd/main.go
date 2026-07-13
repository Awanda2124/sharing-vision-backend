package main

import (
	"log"

	"github.com/awanda/backend-repo/configs"
	"github.com/awanda/backend-repo/internal/handler"
	"github.com/awanda/backend-repo/internal/repository"
	"github.com/awanda/backend-repo/internal/service"
	"github.com/awanda/backend-repo/pkg/database"
	"github.com/awanda/backend-repo/routes"
)

func main() {
	cfg := configs.LoadConfig()

	db, err := database.NewDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	postRepo := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)

	router := routes.SetupRouter(postHandler)

	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
