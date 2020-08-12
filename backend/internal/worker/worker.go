package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt")


func main() {
	d,err:=bcrypt.GenerateFromPassword([]byte("turnbacktime"),6)
    fmt.Println((string)(d),err)
}