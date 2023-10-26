package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	auth_service "gitub.com/regisrex/golang-apis/services/auth"
)

func ValidateJwt() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}
		parsedToken, err := jwt.ParseWithClaims(token, &auth_service.JwtPayload{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_PRIVATE")), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := parsedToken.Claims.(*auth_service.JwtPayload)
		if ok && parsedToken.Valid {
			c.Set("user_id", claims.Id)
			c.Set("user_email", claims.Email)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
		}

	}
}
