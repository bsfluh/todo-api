package api

import (
	"fmt"
	"net/http"
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

// Index
func (h *AuthHandler) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// rigester
func (h *AuthHandler) ShowRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}
func (h *AuthHandler) Register(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	err := h.authService.Register(name, email, password)
	if err != nil {
		c.String(http.StatusInternalServerError, "Cannot register user: %v", err)
		return
	}
	c.Redirect(http.StatusUnauthorized, "/login")
}

// Login
func (h *AuthHandler) ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}
func (h *AuthHandler) Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	user, err := h.authService.Login(email, password)
	if err != nil {
		c.String(401, "Invalid email or password")
		return
	}
	c.SetCookie(
		"user_id",
		fmt.Sprintf("%d", user.ID),
		3600,
		"/",
		"",
		false,
		true,
	)
	c.Redirect(302, "/dashboard")
}
func (h *AuthHandler) Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", gin.H{})
}
