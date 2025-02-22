package middleware

import (
	"avitoshop/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware(signingKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": "Authorization header is required"})
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": "Invalid token format"})
			return
		}

		token := tokenParts[1]

		username, err := jwt.ValidateToken(token, signingKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": "invalid token"})
			return
		}

		c.Set("username", username)
		c.Next()
	}
}
