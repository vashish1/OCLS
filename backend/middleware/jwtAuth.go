package middleware

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	db "github.com/vashish1/OCLS/backend/database"
)

var Secret = []byte(os.Getenv("secret key"))

func GenerateAuthToken(email,name string, user_type int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"type":  user_type,
		"email": email,
		"name": name,
	})

	tokenString, err := token.SignedString([]byte(Secret))
	if err == nil {
		return tokenString, nil
	}
	return "", err
}

func VerifyAuthToken(tokenString string) (bool, map[string]interface{}) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return Secret, nil
	})
	var user_type float64
	var email string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user_type = claims["type"].(float64)
		email = claims["email"].(string)
		_=claims["name"].(string)
	}

	ok, user := db.CheckEmail(email)
	if ok && user["type"].(float64) == user_type {
		return ok, user
	}
	return false, map[string]interface{}{}
}
