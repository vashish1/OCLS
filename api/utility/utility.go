package utility

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/vashish1/OnlineClassPortal/pkg/database/student"
)

var blockKey = os.Getenv("BlockKey")

func GenerateJwtForStudent(email, pass string) (string, error) {
	user, ok := student.Exist(email, pass)
	if ok {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"uid":  user.Uid,
			"name": user.Name,
			"exp":  time.Now().Add(time.Second * 60).Unix(),
		})

		tokenString, err := token.SignedString([]byte(blockKey))
		if err != nil {
			return "", err
		}
		return tokenString, nil
	}
	return "", errors.New("Invalid Credentials")
}

func VerifyJwtForStudent(token string) bool {
	tkn := strings.TrimPrefix(token, "Bearer ")

	token, _ = jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(blockKey), nil
	})

	var m *jwt.MapClaims
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		m["uid"] = claims["uid"].(string)
		m["name"] = claims["name"].(string)
		m["exp"] = claims["exp"].(int64)
	}
	
    ok:=m.VerifyExpiresAt(time.Now().Unix(),true)
    
	return ok
}
