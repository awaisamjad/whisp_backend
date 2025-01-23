// This file is for the handlers related to authentication
package handlers

import (
	// "log"
	"net/http"

	"github.com/awaisamjad/whisp/backend/internal"
	"github.com/awaisamjad/whisp/backend/services"
	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) {
	// Extract form data
	var signUpInfo internal.SignUpRequest

	// Bind JSON data to the struct
	if err := ctx.ShouldBindJSON(&signUpInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, internal.ErrorResponse{ErrorMessage: "Invalid input"})
		return
	}

	userService := service.NewUserService()
	err := userService.SignUp(signUpInfo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, internal.ErrorResponse{ErrorMessage: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)

}

func LogIn(ctx *gin.Context) {

	var logInInfo internal.LogInRequest

	if err := ctx.ShouldBindJSON(&logInInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, internal.ErrorResponse{ErrorMessage: "Invalid input"})
		return
	}

	userService := service.NewUserService()
	logInReturnData, err := userService.LogIn(logInInfo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, internal.ErrorResponse{ErrorMessage: err.Error()})
		return
	}

	jwt_token, err := internal.CreateToken(logInReturnData, 1)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, internal.ErrorResponse{ErrorMessage: err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"username": logInReturnData.Username,
		"id":       logInReturnData.Id,
		"token":    jwt_token,
	})

}

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

	userService := service.NewUserService()

	posts, err := userService.GetPosts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, internal.ErrorResponse{ErrorMessage: err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"posts" : posts,
	})
}