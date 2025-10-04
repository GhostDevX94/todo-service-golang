package pkg

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
	"todo-list/internal/model"
)

func CreateJWTToken(model *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": model.ID,
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
