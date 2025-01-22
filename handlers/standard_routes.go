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