package helpers

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

// Handle error with panic
func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

//Load .env file
func LoadEnv() {
	err := godotenv.Load()
	HandleError(err)
}

// Get jwt secret key from .env
func GetJWTSecret() string {
	LoadEnv()
	secret := os.Getenv("JWT_SECRET")
	return secret
}

// Encrypt password
func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	HandleError(err)

	return string(hash)
}

// Req body validation
func ValidateBody(body interface{}) string {
	validate := validator.New()
	err := validate.Struct(body)
	if err != nil {
		return err.Error()
	}
	return ""
}
