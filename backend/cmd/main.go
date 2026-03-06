package main

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	/*
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
	*/
	r := gin.Default()
	r.LoadHTMLGlob(filepath.Join("..", "templates", "*.html"))
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})
	r.POST("login", func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")
		c.String(200, "Получен логин: email=%s, password=%s", email, password)
	})
	r.Run(":8080")
}
