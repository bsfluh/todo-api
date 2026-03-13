package api

import (
	"todo-Api/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}
func (h *AuthHandler) Register(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	err := h.authService.Register(name, email, password)
	if err != nil {
		c.String(500, "Cannot register user: %v", err)
	}
	c.Redirect(302, "/login")
}
