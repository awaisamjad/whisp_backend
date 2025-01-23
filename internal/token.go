// Code for JWT Tokens

package internal

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(GenerateSecureDailyKey())

//! The time it takes to generate a new key should be the same as the expiry date of the key so its always being made
//! Needs better implementation
func GenerateSecureDailyKey() string {
	currentTime := time.Now()
	year, month, day := currentTime.Date()
	key := fmt.Sprintf("key-%d-%02d-%02d", year, month, day)
	hash := sha256.Sum256([]byte(key))
	return fmt.Sprintf("%x", hash[:])
}
// Duration is in days
func CreateToken(logInReturnData LogInReturn, duration int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": logInReturnData.Username,
			"id":       logInReturnData.Id,
			"exp":      time.Now().Add(time.Second * 24 * time.Duration(duration)).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
