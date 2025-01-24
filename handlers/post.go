package handlers

import (
	"net/http"

	"github.com/awaisamjad/whisp/backend/internal"
	service "github.com/awaisamjad/whisp/backend/services"
	"github.com/gin-gonic/gin"
)

func CreatePost(ctx *gin.Context) {
	var createPostInfo internal.CreatePostRequest

	if err := ctx.ShouldBindJSON(&createPostInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, internal.ErrorResponse{ErrorMessage: "Invalid input"})
		return
	}

	userService := service.NewUserService()
	err := userService.CreatePost(createPostInfo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, internal.ErrorResponse{ErrorMessage: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func GetPosts(ctx *gin.Context) {

	isAuthenticated, exists := ctx.Get("isAuthenticated")
    if !exists {
        isAuthenticated = false // Default to false if not set
    }

	userService := service.NewUserService()

	posts, err := userService.GetPosts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, internal.ErrorResponse{ErrorMessage: err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"isAuthenticated": isAuthenticated,
		"posts" : posts,
	})
}