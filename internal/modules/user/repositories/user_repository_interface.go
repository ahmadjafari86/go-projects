package repositories

import (
	userModels "blog/internal/modules/user/models"
)

type UserRepositoryInterface interface {
	Create(user userModels.User) userModels.User
	FindByEmail(email string) userModels.User
	FindByID(id int) userModels.User
}
