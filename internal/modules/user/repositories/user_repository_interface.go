package repositories

import (
	userModels "blog/internal/modules/user/models"
)

type UserRepositoryInterface interface {
	Create(user userModels.User) userModels.User
}
