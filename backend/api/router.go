package api

import "github.com/gin-gonic/gin"

func SetupRouter(r *gin.Engine, authHandler *AuthHandler, taskHandler *TaskHandler) {
	r.GET("/", authHandler.Index)

	r.GET("/login", authHandler.ShowLogin)
	r.POST("/login", authHandler.Login)

	r.GET("/register", authHandler.ShowRegister)
	r.POST("/register", authHandler.Register)

	auth := r.Group("/")
	auth.Use(AuthMiddleware())
	/*
	   auth.GET("/dashboard",)
	*/

}
