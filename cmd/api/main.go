package main

import (
	"log"
	"net/http"
	"project/internal/application/services"
	"project/internal/core"
	"project/internal/infrastructure/database/postgres"
	"project/internal/infrastructure/repositories"
	"project/internal/presentation/handlers"
	"project/internal/presentation/routers"
)

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

	userRepo, err := repositories.NewUserRepository(db.ConnPool)
	if err != nil {
		log.Fatal(err.Error())
	}

	userService, err := services.NewUserService(userRepo)
	if err != nil {
		log.Fatal(err.Error())
	}

	authHandler, err := handlers.NewAuthHandler(userService)
	if err != nil {
		log.Fatal(err.Error())
	}

	routersOptions := []routers.Option{
		routers.WithAuthRouter(*authHandler),
	}

	router := routers.GetRouter(routersOptions...)
	if err := http.ListenAndServe(cfg.GetAddress(), router); err != nil {
		log.Fatal(err.Error())
	}
}
