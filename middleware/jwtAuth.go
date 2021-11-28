package middleware

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	db "github.com/vashish1/OCLS/database"
)

var Secret = []byte(os.Getenv("secretkey"))

func GenerateAuthToken(email, name string, user_type int) (string, error) {
	fmt.Println(email,"auth generation")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"type":  user_type,
		"email": email,
		"name":  name,
	})

	tokenString, err := token.SignedString([]byte(Secret))
	if err == nil {
		fmt.Println("error while generating auth token",err)
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
		_ = claims["name"].(string)
	}
    fmt.Println(email)
	ok, user := db.CheckEmail(email)
	if ok && user["type"].(float64) == user_type {
		return ok, user
	}
	fmt.Println(ok)
	return false, map[string]interface{}{}
}
