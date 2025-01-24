package middleware

import (
	// "fmt"
	"github.com/awaisamjad/whisp/backend/internal"
	"github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
)

func Auth() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        tokenString := ctx.Request.Header.Get("Authorization")

        // Check if the Authorization header is missing
        if tokenString == "" {
            ctx.Set("isAuthenticated", false)
            ctx.Next()
            return
        }

        // Remove "Bearer " prefix
        if len(tokenString) > len("Bearer ") {
            tokenString = tokenString[len("Bearer "):]
        }

        if tokenString == string(internal.Auth_Token) {
            ctx.Set("isAuthenticated", true)
        } else {
            ctx.Set("isAuthenticated", false)
        }

        ctx.Next()
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
