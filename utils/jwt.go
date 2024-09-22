package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"strings"
	"time"
)

const secretKey = "supersecretkey"

func GenerateToken(email string, userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenString string) (int, error) {
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		log.Println("Error 1", err)
		return -1, errors.New("token could not be parsed")
	}

	isValid := parsedToken.Valid

	if !isValid {
		return -1, errors.New("invalid token")
	}

	claims := parsedToken.Claims.(jwt.MapClaims)

	//email := claims["email"].(string)
	userId := claims["userId"].(float64)

	return int(userId), nil
}
