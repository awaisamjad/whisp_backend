package middleware

import (
	// "fmt"
	"net/http"

	"github.com/awaisamjad/whisp/backend/internal"
	"github.com/gin-gonic/gin"

	// "github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		log.Println("Received Token:", tokenString)

		if len(tokenString) > len("Bearer ") && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		var isAuthenticated bool
        log.Println("This is the Token String :", tokenString)
        log.Println("This is the Internal Auth token:", internal.Auth_Token)
		if tokenString == string(internal.Auth_Token) {
            isAuthenticated = true
			log.Println("User is authenticated")
		} else {
            isAuthenticated = false
			log.Println("User is NOT authenticated")
		}

		// Store authentication status in context
		ctx.Set("isAuthenticated", isAuthenticated)
		ctx.Next()
	}
}

func AttachAuthStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		isAuthenticated, exists := ctx.Get("isAuthenticated")
		if exists {
			ctx.JSON(http.StatusOK, gin.H{
				"isAuthenticated": isAuthenticated,
			})
		}
	}
}

func LogHeaders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for key, values := range ctx.Request.Header {
			log.Printf("%s: %s\n", key, values)
		}
		ctx.Next()
	}
}
