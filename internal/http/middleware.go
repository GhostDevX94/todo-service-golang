package http

import (
	"errors"
	"strings"
	"todo-list/pkg"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		auth := c.GetHeader("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			pkg.ErrorResponse(c, errors.New("unauthorized"), 401)
			c.Abort()
			return
		}
		token := strings.TrimSpace(strings.TrimPrefix(auth, "Bearer "))
		jwtToken, err := pkg.ValidateJWTToken(token)
		if err != nil {
			pkg.ErrorResponse(c, errors.New("unauthorized"), 401)
			c.Abort()
			return
		}

		// MapClaims numeric values decode to float64; cast once and convert
		uidF, ok := jwtToken["uid"].(float64)
		if !ok {
			pkg.ErrorResponse(c, errors.New("unauthorized"), 401)
			c.Abort()
			return
		}
		c.Set("uid", uint(uidF))
		c.Next()
	}
}
