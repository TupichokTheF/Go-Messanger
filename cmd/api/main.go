package main

import (
	"fmt"
	"log"
	"project/internal/core"
	"project/internal/domain/user"
	"project/internal/infrastructure/database/postgres"
	"project/internal/infrastructure/repositories"
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

	testuser, _ := user.CreateUser("TestUser1", "msdad@mail.ru", "1A2w3e4r")
	testRepo := repositories.UserRepository{Pool: db.ConnPool}

	userID, err := testRepo.AddUser(testuser)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(userID)
}
