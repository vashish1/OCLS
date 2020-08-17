package worker

import (
	"golang.org/x/crypto/bcrypt"
)

func Worker(pass string) string {
	d, err := bcrypt.GenerateFromPassword([]byte(pass), 6)
	if err == nil {
		return string(d)
	}
	return ""
}
