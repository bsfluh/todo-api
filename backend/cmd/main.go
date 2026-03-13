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
	userRepo := postgres.NewDBUuserRepository(pool)
	//taskRepo := postgres.NewPoolTaskRepository(pool)

	authService := service.NewAuthService(userRepo)
	//
	authHandler := api.NewAuthHandler(authService)
	r := gin.Default()

	templates := template.Must(template.ParseFS(templatesFS, "templates/*.html"))
	r.SetHTMLTemplate(templates)

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{})
	})
	r.POST("/login", func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")
		c.String(200, "Получен логин: email=%s, password=%s", email, password)
	})
	r.Run(":8080")
}
