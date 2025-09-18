package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jonalphabert/url-shortener-golang/internal/handler"
	"github.com/jonalphabert/url-shortener-golang/internal/logger"
	"github.com/jonalphabert/url-shortener-golang/internal/middleware"
)

func UserRouter(
	h *handler.UserHandler,
	authHandler *handler.AuthHandler,
	urlHandler *handler.UrlHandler,
	log *logger.LoggerType,
) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.RequestLogger(log)) 

	// Public group
	apiAuth := r.Group("/auth")
	authHandler.RegisterRoutes(apiAuth) 

	// Protected group
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware()) 
	{
		h.RegisterRoutes(api)
		urlHandler.RegisterRoutes(api)
	}

	// Public redirect (short url access)
	urlHandler.RegisterRedirectRoutes(r.Group("/"))

	return r
}