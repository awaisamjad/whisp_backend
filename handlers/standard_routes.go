package handlers

import (
	// "log"
	"net/http"

	// "github.com/awaisamjad/whisp/backend/internal"
	"github.com/gin-gonic/gin"
)

func Feed(ctx *gin.Context) {

	ctx.Redirect(http.StatusFound, "/feed")
}

func Settings(ctx *gin.Context) {
	isAuthenticated, exists := ctx.Get("isAuthenticated")

    if !exists {
        isAuthenticated = false
    }

    ctx.JSON(http.StatusOK, gin.H{
        "hello":           "world",
        "isAuthenticated": isAuthenticated,
    })
}