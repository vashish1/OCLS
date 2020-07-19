package utility

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/vashish1/OnlineClassPortal/pkg/database/student"
	"github.com/vashish1/OnlineClassPortal/pkg/database/teacher"
	"github.com/vashish1/OnlineClassPortal/pkg/helpers"
)

var blockKey = os.Getenv("BlockKey")

func GenerateJwtForStudent(email, pass string) (string, error) {
	user, ok := student.Exist(email, pass)
	if ok {
		ok := helpers.ValidatePass([]byte(user.PassHash), pass)
		if ok {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"uid":  user.Uid,
				"name": user.Name,
				"exp":  time.Now().Add(time.Hour * 72).Unix(),
			})

			tokenString, err := token.SignedString([]byte(blockKey))
			if err != nil {
				fmt.Println(err)
				return "", err
			}
			return tokenString, nil
		}

	}
	return "", errors.New("Invalid Credentials")
}

func VerifyJwt(t string) bool {
	tkn := strings.TrimPrefix(t, "Bearer ")

	token, _ := jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(blockKey), nil
	})

	var uid, name string
	var exp float64
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uid = claims["uid"].(string)
		name = claims["name"].(string)
		exp = claims["exp"].(float64)
	}

	m := &jwt.MapClaims{
		"uid":  uid,
		"name": name,
		"exp":  exp,
	}
	ok := m.VerifyExpiresAt(time.Now().Unix(), true)
	if ok {
		return true
	}
	return false
}

func GenerateJwtForTeacher(email, pass string) (string, error) {
	user, ok := teacher.Exist(email, pass)
	if ok {
		ok := helpers.ValidatePass([]byte(user.PassHash), pass)
		if ok {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"uid":  user.Uid,
				"name": user.Name,
				"exp": time.Now().Add(time.Hour * 120).Unix(),
			})

			tokenString, err := token.SignedString([]byte(blockKey))
			if err != nil {
				fmt.Println(err)
				return "", err
			}
			return tokenString, nil
		}

	}
	return "", errors.New("Invalid Credentials")
}
