package main

import (
	"fmt"
	"log"
	"net/http"
	"project/internal/application/services"
	"project/internal/core"
	"project/internal/infrastructure/database/postgres"
	"project/internal/infrastructure/repositories"

	//"project/internal/infrastructure/repositories"
	"project/internal/presentation/handlers"
	"project/internal/presentation/routers"
)

// @title    Go Messenger API
// @version  1.0
// @host     localhost:8080
// @BasePath /api/v1
func main() {
	start()
}

func start() {
	cfg := core.LoadConfig()

	db, err := postgres.OpenConnection(cfg.DataBaseConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseConnection()

	userRepo := repositories.NewUserRepository(db.ConnPool)

	userService := services.NewUserService(userRepo)

	authHandler := handlers.NewAuthHandler(userService)

	routersOptions := []routers.Option{
		routers.WithAuthRouter(*authHandler),
	}

	if cfg.Swagger {
		fmt.Println("URAAA")
		routersOptions = append(routersOptions, routers.WithSwagger())
	}

	router := routers.GetRouter(routersOptions...)
	if err := http.ListenAndServe(cfg.GetAddress(), router); err != nil {
		log.Fatal(err.Error())
	}
}
