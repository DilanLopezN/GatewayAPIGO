package main

import (
	"database/sql"
	"fmt"
	"gateway/internals/http/server"
	"gateway/internals/repository"
	"gateway/internals/services"
	"log"
	"os"

	"github.com/joho/godotenv"
)


func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

connStr := fmt.Sprintf(
	"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
	getEnv("DB_HOST", "db"),
	getEnv("DB_PORT", "5432"),
	getEnv("DB_USER", "postgres"),
	getEnv("DB_PASSWORD", "postgres"),
	getEnv("DB_NAME", "postgres"),
	getEnv("DB_SSLMODE", "disable"),
)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	accountRepository := repository.NewAccountRepository(db)
	accountService := services.NewAccountService(accountRepository)

	port := getEnv("PORT", "8080")

	srv := server.NewServer(accountService, port)

	srv.ConfigureRoutes()

	if err := srv.Start(); 
	err != nil {
		log.Fatal("Error starting server: ", err)
	}


}
