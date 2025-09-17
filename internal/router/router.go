package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jonalphabert/url-shortener-golang/internal/handler"
	"github.com/jonalphabert/url-shortener-golang/internal/logger"
)

// func UserRouter(h *handler.UserHandler, log *logger.LoggerType) *gin.Engine {
//     r := gin.New()
//     r.Use(gin.Recovery())
//     r.Use(RequestLogger(log)) // custom middleware

//     api := r.Group("/api")
//     h.RegisterRoutes(api)
//     u.RegisterRoutes(api)

//     return r
// }

func UserRouter(h *handler.UserHandler,  log *logger.LoggerType) *gin.Engine {
    r := gin.New()
    r.Use(gin.Recovery())
    r.Use(RequestLogger(log)) // custom middleware

    api := r.Group("/api")
    h.RegisterRoutes(api)

    return r
}

func RequestLogger(log *logger.LoggerType) gin.HandlerFunc {
    return func(c *gin.Context) {
        path := c.Request.URL.Path
        method := c.Request.Method
        log.Infof("Started %s %s", method, path)

        c.Next()

        status := c.Writer.Status()
        log.Infof("Completed %d %s %s", status, method, path)
    }
}
