package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jonalphabert/url-shortener-golang/internal/logger"
)

func RequestLogger(log *logger.LoggerType) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next() // proses request

		latency := time.Since(start)
		log.Info(
			"Request completed",
			"; method: ", c.Request.Method,
			"; path: ", c.Request.URL.Path,
			"; status: ", c.Writer.Status(),
			"; latency: ", latency.String(),
		)
	}
}