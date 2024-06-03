package tools

import "golang.org/x/crypto/bcrypt"

// The HashPassword function takes a password string, generates a hashed version of it using bcrypt
// with a cost factor of 14, and returns the hashed password as a string along with any error
// encountered.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
