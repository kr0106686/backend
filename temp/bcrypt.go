package temp

import (
	"golang.org/x/crypto/bcrypt"
)

func Bcrypt(pass string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pass), 10)
	return string(hash)
}

func CompareHashAndPassword(hash, pass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err != nil {
		return err
	}
	return nil
}
