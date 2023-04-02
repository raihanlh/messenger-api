package utils

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gitlab.com/raihanlh/messenger-api/config"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	jwt.RegisteredClaims
	Email  string `json:"email"`
	UserID string `json:"id"`
	Exp    int64  `json:"exp"`
}

func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return bytes, err
}

func GenerateToken(email string, id string, exp time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Email:  email,
		UserID: id,
		Exp:    exp.Unix(),
	})

	conf := config.New()

	t, err := token.SignedString([]byte(conf.Secret))
	if err != nil {
		log.Println("failed to create token", err.Error())
		return "", err
	}

	return t, nil
}
