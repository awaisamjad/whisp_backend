// This file is for creating services to be used related to authentication
package service

import (
	"fmt"
	log "github.com/sirupsen/logrus"

	"github.com/awaisamjad/whisp/backend/internal"
	"github.com/awaisamjad/whisp/backend/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		repo: repositories.NewUserRepository(),
	}
}

func (u *UserService) SignUp(signUpInfo internal.SignUpRequest) error {

	first_name := signUpInfo.FirstName
	last_name := signUpInfo.LastName
	username := signUpInfo.Username
	email := signUpInfo.Email
	password := signUpInfo.Password
	confirm_password := signUpInfo.ConfirmPassword

	if password != confirm_password {
		return fmt.Errorf("Passwords do not match")
	}

	// ? Validating inputs
	if !internal.IsNameValid(first_name) || !internal.IsNameValid(last_name) {
		return fmt.Errorf("Invalid First or Last name. Ensure they are between 2 and 50 characters long, contain no special characters, and have no spaces.")
	}

	if !internal.IsUsernameValid(username) {
		return fmt.Errorf("Invalid Username. Ensure it is between 3 and 20 characters long, and has no spaces.")
	}

	if !internal.IsEmailValid(email) {
		return fmt.Errorf("Invalid email")
	}
	if !internal.IsPasswordValid(password) {
		return fmt.Errorf(`Password does not meet the required criteria. 
		Ensure it is at least 8 characters long, includes at least one uppercase letter, 
		one lowercase letter, one digit, and one special character (e.g., !@#$^&*...).`)
	}

	password, err := internal.HashPassword(password)
	if err != nil {
		log.Error("Error hashing password")
		return fmt.Errorf("Internal server error")
	}
	signUpInfo.Password = password
	return u.repo.CreateUser(signUpInfo)
}

func (u *UserService) LogIn(logInInfo internal.LogInRequest) (internal.LogInReturn, error) {

	return u.repo.LogInUser(logInInfo)
}

func (u *UserService) CreatePost(createPostInfo internal.CreatePostRequest) error {
	if createPostInfo.Content == "" {
		return fmt.Errorf("The post content is empty")
	}

	return u.repo.CreatePost(createPostInfo)
}

func (u *UserService) GetPosts() ([]internal.Post, error) {
	return u.repo.GetPosts()
}

func (u *UserService) GetPostByID(id string) (internal.Post, error) {
	return u.repo.GetPostByID(id)
}
