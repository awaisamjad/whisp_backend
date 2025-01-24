// Code for JWT Tokens

package internal

import (
	"crypto/rand"
	"encoding/base64"
	// "fmt"
	log "github.com/sirupsen/logrus"
	// "os"
	// "time"
	// "github.com/golang-jwt/jwt/v5"
)

// var Auth_Token = []byte("hellohellueohellohellohellohellos")
var Auth_Token = (GenerateSecureDailyKey())

// ! The time it takes to generate a new key should be the same as the expiry date of the key so its always being made
// ! Needs better implementation

func GenerateSecureDailyKey() string {
    // Create a byte slice to hold the random bytes
    token := make([]byte, 32) // 32 bytes = 256 bits

    // Read random bytes from the crypto/rand package
    _, err := rand.Read(token)
    if err != nil {
        // return "", fmt.Errorf("failed to generate random bytes: %w", err)
		log.Error(err)
    }

    // Encode the random bytes to a base64 string
    // Base64 encoding will include non-alphanumeric characters, so we will replace them
    base64Token := base64.RawURLEncoding.EncodeToString(token)

    // Remove non-alphanumeric characters (if needed)
    alphanumericHash := ""
    for _, char := range base64Token {
        if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
            alphanumericHash += string(char)
        }
    }

    // Ensure the length of the key is as expected (e.g., 32 characters)
    if len(alphanumericHash) < 32 {
        // return "", fmt.Errorf("generated key is too short: %d characters", len(alphanumericHash))
		log.Error(err)
    }

    return alphanumericHash[:32] // Return the first 32 characters
}

// Duration is in days
// func CreateToken(logInReturnData LogInReturn, duration int) (string, error) {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
// 		jwt.MapClaims{
// 			"username": logInReturnData.Username,
// 			"id":       logInReturnData.User_Id,
// 			"exp":      time.Now().Add(time.Hour * 24 * time.Duration(duration)).Unix(),
// 		})

// 	tokenString, err := token.SignedString(JWTSecret)
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }

// func VerifyToken(tokenString string) error {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return JWTSecret, nil
// 	})

// 	if err != nil {
// 		return err
// 	}

// 	if !token.Valid {
// 		return fmt.Errorf("invalid token")
// 	}

// 	return nil
// }
