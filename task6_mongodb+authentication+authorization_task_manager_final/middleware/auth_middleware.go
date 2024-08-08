package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"task_manager/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
) 
 
var jwtKey = []byte("my_secret_key") 
 
type Claims struct {
	 UserID primitive.ObjectID `json:"userId"`
	 Username string `json:"username"`
	 Role string `json:"role"`
	 jwt.StandardClaims
}
 
func GenerateJWT(user models.User)(string,error){ 
 expirationTime := time.Now().Add(24 * time.Hour) 
 claims := &Claims{ 
        UserID:   user.ID, 
        Username: user.Username, 
        Role:     user.Role, 
        StandardClaims: jwt.StandardClaims{ 
            ExpiresAt: expirationTime.Unix(), 
        }, 
    } 
 
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) 
    tokenString, err := token.SignedString(jwtKey) 
    if err != nil { 
        return "", err 
    } 
    return tokenString, nil 
} 
 
func ValidateJWT(tokenString string) (*Claims, error) { 
    claims := &Claims{} 
 
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) { 
        return jwtKey, nil 
    }) 
    if err != nil { 
        if err == jwt.ErrSignatureInvalid { 
            return nil, errors.New("invalid signature") 
        } 
        return nil, err 
    } 
    if !token.Valid { 
        return nil, errors.New("validatejwt func invalid token") 
    } 
    return claims, nil 
} 
 
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
            c.Abort()
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")

        user, err := ValidateJWT(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
            c.Abort()
            return
        }

        c.Set("user", user)
        c.Next()
    }
}

func AdminMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        user, exists := c.Get("user")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
            c.Abort()
            return
        }

        claims, ok := user.(*Claims)
        if !ok {
            c.JSON(http.StatusForbidden, gin.H{"error": "forbidden - invalid claims"})
            c.Abort()
            return
        }

        fmt.Println("User Claims:", claims.Role) // Debugging line

        if claims.Role != "admin" {
            c.JSON(http.StatusForbidden, gin.H{"error": "forbidden - not an admin"})
            c.Abort()
            return
        }

        c.Next()
    }
}
