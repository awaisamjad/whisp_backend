// This file is for the handlers related to authentication
package handlers

import (
	log "github.com/sirupsen/logrus"
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

	log.Println(string(internal.Auth_Token))
	ctx.JSON(http.StatusOK, gin.H{
		"username":   logInReturnData.Username,
		"user_id":    logInReturnData.User_Id,
		"auth_token": internal.Auth_Token,
		"avatar":     logInReturnData.Avatar,
	})
	log.Println(string(internal.Auth_Token))

}
