package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jonalphabert/url-shortener-golang/internal/db"
	"github.com/jonalphabert/url-shortener-golang/internal/handler"
	"github.com/jonalphabert/url-shortener-golang/internal/logger"
	"github.com/jonalphabert/url-shortener-golang/internal/models"
	"github.com/jonalphabert/url-shortener-golang/internal/repository"
	"github.com/jonalphabert/url-shortener-golang/internal/router"
	"github.com/jonalphabert/url-shortener-golang/internal/service"

	"github.com/joho/godotenv"
)

func main() {
    log := logger.New() // pakai logger ada
    log.Info("Starting app")

    // load env
    if err := godotenv.Load(".env"); err != nil {
        log.Fatal(err)
    }

    // Koneksi ke DB
    db, err := db.Connect(os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatal(err)
    } 

    log.Info("Connected to DB")
    

    // Auto migrate
    if err := db.AutoMigrate(&models.User{}, &models.Url{}); err != nil {
        log.Fatal(err)
    }

    log.Info("Migrated DB has been successfully")

    // repository (in-memory sekarang)
    userRepo := repository.NewUserRepository(db)
    urlRepo := repository.NewUrlRepository(db)

    // service
    userSvc := service.NewUserService(userRepo, log)
    authSvc := service.NewAuthService(userRepo, log)
    urlSvc := service.NewUrlService(urlRepo, log)

    // handler
    userHandler := handler.NewUserHandler(userSvc, log)
    authHandler := handler.NewAuthHandler(authSvc, log)
    urlHandler := handler.NewUrlHandler(urlSvc, log)

    // router
    r := router.UserRouter(userHandler, authHandler, urlHandler, log)

    r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

    r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })

    r.Run(":8080")
}
