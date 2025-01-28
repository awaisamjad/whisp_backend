package main

import (
	// "time"
	"fmt"
	"path/filepath"
	"runtime"

	// "log"
	"net/http"

	log "github.com/sirupsen/logrus"

	// "net/http"
	"os"

	"github.com/awaisamjad/whisp/backend/db"
	"github.com/awaisamjad/whisp/backend/handlers"
	"github.com/awaisamjad/whisp/backend/middleware"

	// "github.com/awaisamjad/whisp/backend/internal"
	"github.com/joho/godotenv"
	// "github.com/gorilla/sessions"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)

	log.SetFormatter(&log.TextFormatter{
		PadLevelText: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			_, b, _, _ := runtime.Caller(0)
			basepath := filepath.Dir(b)
			rel, err := filepath.Rel(basepath, f.File)
			if err != nil {
				log.Error("Couldn't determine file path\n", err)
			}
			return "", fmt.Sprintf("%s:%d", rel, f.Line)
		},
	})
	log.SetReportCaller(true)

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	db, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	r := gin.Default()
	r.Use(middleware.Auth())

	FRONTEND_LOCAL_URL := os.Getenv("FRONTEND_LOCAL_URL")
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{FRONTEND_LOCAL_URL}, // Ensure this matches your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	standardRoutes := r.Group("/")
	standardRoutes.GET("/", handlers.Feed)

	auth := r.Group("/auth")
	auth.POST("/signup", handlers.SignUp)
	auth.POST("/login", middleware.LogHeaders(), handlers.LogIn)

	r.POST("/create-post", handlers.CreatePost)
	r.GET("/posts", handlers.GetPosts)
	r.GET("/posts/:id", handlers.GetPostByID)
	r.GET("/test", func(ctx *gin.Context) {

	})

	r.GET("/user/:username", func(ctx *gin.Context) {
		username := ctx.Param("username")
		query := `SELECT * FROM posts WHERE username=?, (username)`
		posts, err := db.Exec(query, &username)
		if err != nil {
			log.Error("Failed to insert user post in posts table")
			// return fmt.Errorf("Failed to create post")
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"ErrorMessage": "VERY big ERROR",
			})
			return
		}
		log.Println(posts)
	})

	port := os.Getenv("PORT")
	if port == "" {
		log.Error("No port specified in .env, using default 8000")
		port = "8000"
	}

	// Run the server
	r.Run("0.0.0.0:" + port)
}
