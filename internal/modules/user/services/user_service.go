package services

import (
	userModels "blog/internal/modules/user/models"
	userRepositories "blog/internal/modules/user/repositories"
	"blog/internal/modules/user/requests/auth"
	userResponses "blog/internal/modules/user/responses"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository userRepositories.UserRepositoryInterface
}

func New() *UserService {
	return &UserService{
		userRepository: userRepositories.New(),
	}
}

func (userService *UserService) Create(request auth.RegisterRequest) (userResponses.User, error) {
	var response userResponses.User
	var user userModels.User

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	if err != nil {
		return response, errors.New("error in hashing password")
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Password = string(hashedPassword)

	newUser := userService.userRepository.Create(user)

	if newUser.ID == 0 {
		return response, errors.New("error on user creation process")
	}
	return userResponses.ToUser(newUser), nil
}

func (userService *UserService) CheckUserExists(email string) bool {
	user := userService.userRepository.FindByEmail(email)
	return user.ID != 0
}

func (userService *UserService) HandleUserLogin(request auth.LoginRequest) (userResponses.User, error) {
	var response userResponses.User
	existsUser := userService.userRepository.FindByEmail(request.Email)
	if existsUser.ID == 0 {
		return response, errors.New("user not exists")
	}
	err := bcrypt.CompareHashAndPassword([]byte(existsUser.Password), []byte(request.Password))
	if err != nil {
		return userResponses.ToUser(existsUser), errors.New("incorrect password")
	}

	return userResponses.ToUser(existsUser), nil
}
