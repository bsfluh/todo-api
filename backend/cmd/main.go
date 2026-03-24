package main

import (
	"embed"
	"html/template"
	"log"
	"todo-Api/api"
	"todo-Api/config"
	"todo-Api/db"
	"todo-Api/repository/postgres"
	"todo-Api/service"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

//go:embed templates/*.html
var templatesFS embed.FS

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
	// Repositories
	userRepo := postgres.NewPoolUserRepository(pool)
	taskRepo := postgres.NewPoolTaskRepository(pool)
	// Services
	authService := service.NewAuthService(userRepo)
	taskService := service.NewTaskService(taskRepo)
	// Handlers
	taskHandler := api.NewTaskHandler(taskService)
	authHandler := api.NewAuthHandler(authService)

	r := gin.Default()

	templates := template.Must(template.ParseFS(templatesFS, "templates/*.html"))
	r.SetHTMLTemplate(templates)
	api.SetupRouter(r, authHandler, taskHandler)

	r.Run(":8080")
}
