package helper

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	saltStr := os.Getenv("BCRYPT_SALT")
	salt, err := strconv.Atoi(saltStr)
	ErrorIfPanic(err)
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	ErrorIfPanic(err)
	return string(hashPassword)
}

func CompareHashPassword(passwordHash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}
