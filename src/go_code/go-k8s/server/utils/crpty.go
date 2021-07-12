package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Encode(pwd string) string {
	// TODO
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("init encode password failed.")
	}
	encodePwd := string(hash)
	return encodePwd
}
