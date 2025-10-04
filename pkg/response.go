package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type OperationSuccess struct {
	Status bool
	Data   interface{}
}

type TokenR struct {
	Status bool
	Token  interface{}
	Data   interface{}
}

func CreatedResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, OperationSuccess{Status: true, Data: data})
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, OperationSuccess{Status: true, Data: data})
}

func TokenResponse(c *gin.Context, data interface{}, token string) {
	c.JSON(http.StatusOK, TokenR{
		Status: true,
		Data:   data,
		Token:  token,
	})
}

func ErrorResponse(c *gin.Context, error error, code int) {
	c.JSON(code, gin.H{"error": error.Error()})
}
