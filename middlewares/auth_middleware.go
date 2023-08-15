package middlewares

import (
	"net/http"

	"github.com/aryansh13/go-restapi-gin/helpers"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
			c.Abort()
			return
		}

		token, err := helpers.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Di sini Anda bisa melakukan validasi tambahan,
		// seperti memeriksa apakah pengguna masih ada dalam basis data, dll.

		c.Set("userID", token.UserID) // Menyimpan ID pengguna dalam konteks
		c.Next()                      // Melanjutkan ke handler
	}
}
