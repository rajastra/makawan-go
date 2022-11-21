package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(password string) (string, error) {
	passHashByte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Println("Generate Password Error :", err)
		return "", err
	}
	return string(passHashByte), nil
}

func ValidatePassword(passHashdb string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(passHashdb), []byte(password)) == nil
}
