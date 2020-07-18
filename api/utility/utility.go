package utility

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJwtForStudent(email, pass string) (string, error) {
	user, ok := student.Exist(email, pass)
	if ok {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"uid": user.Uid,
			"":    user.Name,
		})

		tokenString, err := token.SignedString([]byte("thesportspritslies"))
		if err != nil {
			return "", err
		}
		return tokenString, nil
	}
	return "", errors.New("Invalid Credentials")
}
