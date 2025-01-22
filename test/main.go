package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	// "os"
	// "github.com/joho/godotenv"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	fmt.Println("Error loading .env file")
	// }
	// os.Setenv("SESSION_KEY", "Bye")
	// key := os.Getenv("SESSION_KEY")
	// os.Setenv("SESSION_KEY", "Hello")
	// key2 := os.Getenv("SESSION_KEY")

	// fmt.Println(key, key2)
	fmt.Println(generateRandomKey(32))
}

func generateRandomKey(length int) string {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatal("Failed to generate random key:", err)
	}
	return base64.StdEncoding.EncodeToString(key)
}
