package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/widcha/openidea-marketplace/configs"
	"github.com/widcha/openidea-marketplace/internal/pkg/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientTokenBearer := c.Request.Header.Get("Authorization")
		clientToken := strings.Split(clientTokenBearer, " ")[1]
		if clientToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "unauthorized",
				"data":    nil,
			})
			return
		}

		parsedToken, err := jwt.ParseWithClaims(clientToken, &token.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(configs.Get().JwtSecret), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "unauthorized",
				"data":    nil,
			})
			return
		}

		if !parsedToken.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "unauthorized",
				"data":    nil,
			})
			return
		}

		uc, isUserClaim := parsedToken.Claims.(*token.UserClaims)
		if uc.UserID == "" || !isUserClaim {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "unauthorized",
				"data":    nil,
			})
			return
		}

		c.Set("user", parsedToken.Claims)
		c.Next()
	}
}
