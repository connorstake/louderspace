// Intializes the application backend and starts the server
package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"louderspace/config"
	"louderspace/internal/api"
	"louderspace/internal/repositories"
	"louderspace/internal/services"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	userStorage := repositories.NewUserDatabase(db)
	userService := services.NewUserService(userStorage)
	userAPI := api.NewUserAPI(userService)

	http.HandleFunc("/register", userAPI.Register)
	http.HandleFunc("/login", userAPI.Login)

	log.Println("server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
