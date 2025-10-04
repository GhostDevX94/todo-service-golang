package http

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"todo-list/pkg"
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

		token := c.GetHeader("Authorization")

		if token == "" {
			pkg.ErrorResponse(c, errors.New("unauthorized"), 401)
			c.Abort()
			return
		}

		BearerSlice := bytes.Split([]byte(token), []byte(" "))
		Token := string(BearerSlice[1])

		jwtToken, err := pkg.ValidateJWTToken(Token)
		if err != nil {
			pkg.ErrorResponse(c, errors.New("unauthorized"), 401)
			c.Abort()
			return
		}

		c.Set("uid", jwtToken["uid"])

		c.Next()
	}
}
