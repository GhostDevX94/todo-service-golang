package http

import (
	"os"
	"strings"
	"time"
	"todo-list/internal/errors"
	"todo-list/pkg"

	"github.com/gin-gonic/gin"
)

func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		duration := time.Since(start)
		statusCode := c.Writer.Status()

		pkg.Logger.Info().
			Str("method", method).
			Str("path", path).
			Int("status", statusCode).
			Dur("duration", duration).
			Str("ip", c.ClientIP()).
			Msg("Request processed")
	}
}

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				pkg.Logger.Error().
					Interface("error", err).
					Str("path", c.Request.URL.Path).
					Msg("Panic recovered")

				c.JSON(500, gin.H{
					"error":   "Internal server error",
					"message": "An unexpected error occurred",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedOrigins := os.Getenv("CORS_ALLOWED_ORIGINS")
		if allowedOrigins == "" {
			allowedOrigins = "*"
		}

		c.Header("Access-Control-Allow-Origin", allowedOrigins)
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length")
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
			appErr := errors.Unauthorized("Missing or invalid authorization header")
			c.JSON(appErr.Code, gin.H{
				"error":   appErr.Err.Error(),
				"message": appErr.Message,
			})
			c.Abort()
			return
		}

		token := strings.TrimSpace(strings.TrimPrefix(auth, "Bearer "))
		jwtToken, err := pkg.ValidateJWTToken(token)
		if err != nil {
			appErr := errors.Unauthorized("Invalid or expired token")
			pkg.Logger.Warn().
				Err(err).
				Str("ip", c.ClientIP()).
				Msg("Authentication failed")

			c.JSON(appErr.Code, gin.H{
				"error":   appErr.Err.Error(),
				"message": appErr.Message,
			})
			c.Abort()
			return
		}

		uidF, ok := jwtToken["uid"].(float64)
		if !ok {
			appErr := errors.Unauthorized("Invalid token claims")
			c.JSON(appErr.Code, gin.H{
				"error":   appErr.Err.Error(),
				"message": appErr.Message,
			})
			c.Abort()
			return
		}

		c.Set("uid", uint(uidF))
		c.Next()
	}
}
