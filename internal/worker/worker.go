package worker

import (
	"fmt"

	"golang.org/x/crypto/bcrypt")


func worker() {
	d,err:=bcrypt.GenerateFromPassword([]byte("turnbacktime"),6)
    fmt.Println((string)(d),err)
}