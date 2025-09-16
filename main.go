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
    urlRepo := repository.NewInMemoryUrlRepo()

    // seed data (opsional)
    userRepo.Create(&models.User{Name: "Jonathan"})
    userRepo.Create(&models.User{Name: "Seth"})
    userRepo.Create(&models.User{Name: "John"})

    urlRepo.Create(&models.Url{ShortUrl: "backtracking", LongUrl: "https://chatgpt.com/c/68c8be98-2460-8322-97d2-39028cce4be5"})

    // service
    userSvc := service.NewUserService(userRepo, log)
    urlSvc := service.NewUrlService(urlRepo, log)

    // handler
    userHandler := handler.NewUserHandler(userSvc, log)
    urlHandler := handler.NewUrlHandler(urlSvc, log)

    // router
    r := router.UserRouter(userHandler, urlHandler, log)

    r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

    r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })

    r.Run(":8080")
}
