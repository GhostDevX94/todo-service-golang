package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type OperationSuccess struct {
	Status bool
	Data   interface{}
}

func CreatedResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, OperationSuccess{Status: true, Data: data})
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, OperationSuccess{Status: true, Data: data})
}

func ErrorResponse(c *gin.Context, error error, code int) {
	c.JSON(code, gin.H{"error": error.Error()})
}
