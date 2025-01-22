// This file is store utility/useful functions to be used throughout the package
package internal

import (
	"regexp"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

// ? Password Rules
// Minimum length : 8 Characters
// Must include at least one `UPPERCASE` letter
// Must include at least one `lowercase` letter
// Must include at least one digit
// Must include at least one special character (e.g., !@#$%^&*...)
func IsPasswordValid(password string) bool {
	if len(password) < 8 {
		return false
	} 

	var hasUpper, hasLower, hasDigit, hasSpecial bool

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasDigit && hasSpecial
}

// Must be between 3 and 20 characters long
// No spaces or special characters
func IsUsernameValid(username string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9_]{3,20}$`)
	return re.MatchString(username)
}

// TODO send email and get user to verify
func IsEmailValid(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// Only alphabetic characters
// No spaces
// Between 2 and 50 Characters
func IsNameValid(name string) bool {
	re := regexp.MustCompile(`^[a-zA-Z]{2,50}$`)
	return re.MatchString(name)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares a hashed password with a plain text password
func CheckPasswordAgainstPasswordHash(password, hash_password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash_password), []byte(password))
	return err == nil
}
