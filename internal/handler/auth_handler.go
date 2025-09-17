package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonalphabert/url-shortener-golang/internal/logger"
	"github.com/jonalphabert/url-shortener-golang/internal/service"
)

type AuthHandler struct {
	svc *service.AuthServices
	log *logger.LoggerType
}

func NewAuthHandler(svc *service.AuthServices, log *logger.LoggerType) *AuthHandler {
	return &AuthHandler{svc: svc, log: log}
}

func (h *AuthHandler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/login", h.Login)
	rg.POST("/register", h.Register)
}

func (h *AuthHandler) Login(c *gin.Context) {
	h.log.Info("Login called")

	var body struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	auth, err := h.svc.Login(body.Username, body.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": auth.Token, "id": auth.ID, "username": auth.Username, "status": "Login Success"})
}

func (h *AuthHandler) Register(c *gin.Context) {
	h.log.Info("Register called")

	var body struct {
        Username     string `json:"username" binding:"required"`
        Password 	string `json:"password" binding:"required"`
    }

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	user, err := h.svc.Register(body.Username, body.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "Registration Success",
		"user": gin.H{
			"id": user.ID,
			"username": user.Username,
		},
	})
}