package helpers

import "golang.org/x/crypto/bcrypt"

func EncodePass(pass string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 20)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func ValidatePass(hash []byte,pass string) bool {
	 err:=bcrypt.CompareHashAndPassword(hash,[]byte(pass))
	 if err!=nil{
		 return false
	 }
	 return true
}
