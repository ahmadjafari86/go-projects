package services

import (
	"blog/internal/modules/user/requests/auth"
	userResponses "blog/internal/modules/user/responses"
)

type UserServiceInterface interface {
	CheckUserExists(email string) bool
	Create(request auth.RegisterRequest) (userResponses.User, error)
	HandleUserLogin(request auth.LoginRequest) (userResponses.User, error)
}
