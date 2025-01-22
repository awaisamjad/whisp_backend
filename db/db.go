package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// "github.com/awaisamjad/whisp/backend/internal"
	// "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func ConnectDB() (*sql.DB, error) {
	
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_PORT := os.Getenv("DB_PORT")

	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp(127.0.0.1:" + DB_PORT + ")/whisp_db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Verify the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connected to MySQL!")
	return db, nil
}

