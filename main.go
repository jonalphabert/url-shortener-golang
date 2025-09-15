package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jonalphabert/url-shortener-golang/handler"
	"github.com/jonalphabert/url-shortener-golang/logger"
	"github.com/jonalphabert/url-shortener-golang/models"
	"github.com/jonalphabert/url-shortener-golang/repository"
	"github.com/jonalphabert/url-shortener-golang/router"
	"github.com/jonalphabert/url-shortener-golang/service"
)

func main() {
    log := logger.New() // pakai logger ada
    log.Info("Starting app")

    // repository (in-memory sekarang)
    userRepo := repository.NewInMemoryUserRepo()

    // seed data (opsional)
    userRepo.Create(&models.User{Name: "Jonathan"})
    userRepo.Create(&models.User{Name: "Seth"})
    userRepo.Create(&models.User{Name: "John"})

    // service
    userSvc := service.NewUserService(userRepo, log)

    // handler
    userHandler := handler.NewUserHandler(userSvc, log)

    // router
    r := router.UserRouter(userHandler, log)

    r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

    r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })

    r.Run(":8080")
}
