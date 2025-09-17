package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jonalphabert/url-shortener-golang/internal/logger"
	"github.com/jonalphabert/url-shortener-golang/internal/service"
)

type UserHandler struct {
    svc *service.UserService
    log *logger.LoggerType
}

func NewUserHandler(svc *service.UserService, log *logger.LoggerType) *UserHandler {
    return &UserHandler{svc: svc, log: log}
}

func (h *UserHandler) RegisterRoutes(rg *gin.RouterGroup) {
    rg.GET("/users", h.GetAll)
    rg.GET("/user/:id", h.GetByID)
    rg.POST("/user", h.Create)
    rg.DELETE("/user/:id", h.Delete)
}

func (h *UserHandler) GetAll(c *gin.Context) {
    h.log.Info("GetAllUsers called")
    users, err := h.svc.GetAllUsers()
    if err != nil {
        h.log.WithError(err).Error("GetAllUsers failed")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"users": users})
}

func (h *UserHandler) GetByID(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    user, err := h.svc.GetUser(id)
    if err != nil {
        h.log.WithField("id", id).Warn("User not found")
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *UserHandler) Create(c *gin.Context) {
    var body struct {
        Name     string `json:"name" binding:"required"`
        Password string `json:"password" binding:"required"`
    }
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
        return
    }
    user, err := h.svc.CreateUser(body.Name, body.Password)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (h *UserHandler) Delete(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    u, err := h.svc.DeleteUser(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"user": u})
}
