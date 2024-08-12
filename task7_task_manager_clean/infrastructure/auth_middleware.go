package infrastructure

import (
	"net/http"
	"strings"
	// "tskmgr/domain"
	"github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt"
	// "golang.org/x/crypto/bcrypt"
    // "errors"
)




func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		c.Abort()
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Set("claim", *claims)
	c.Next()
}
