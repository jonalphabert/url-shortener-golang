package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jonalphabert/url-shortener-golang/internal/utils"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
		// Ambil "Authorization: Bearer <token>"
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization format"})
			c.Abort()
			return
		}

		// Parse JWT
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// Inject ke context biar bisa dipakai di handler
		c.Set("userID", claims.UserID)

		c.Next()
	}
}
