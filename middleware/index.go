package middleware

import (
	"net/http"
	"strings"
	"todolist-api/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware memastikan bahwa request memiliki token JWT yang valid
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil token dari header Authorization
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		// Hapus "Bearer " dari token jika ada
		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) == 2 {
			tokenString = tokenParts[1]
		}

		// Validasi token
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Simpan userID ke context
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
