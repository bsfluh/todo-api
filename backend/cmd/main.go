package main

import (
	"log"
	"todo-Api/config"
	"todo-Api/db"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("error config: %v", err)
	}
	pool, err := db.NewPostgresPool(cfg)
	if err != nil {
		log.Fatalf("error db: %v", err)
	}
	defer pool.Close()
	//дбшкка

	r := gin.Default()
	r.Run(":8080")
}
