package helpers

import (
	userRepositories "blog/internal/modules/user/repositories"
	userResponses "blog/internal/modules/user/responses"
	"blog/pkg/sessions"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) userResponses.User {
	var response userResponses.User
	var userRepo = userRepositories.New()
	authID := sessions.Get(c, "auth")
	userID, _ := strconv.Atoi(authID)
	user := userRepo.FindByID(userID)
	if user.ID == 0 {
		return response
	}

	return userResponses.ToUser(user)
}
