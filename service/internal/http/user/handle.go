package user

import (
	"net/http"
	"user_balance/service/internal/logger"
	userService "user_balance/service/internal/service/user"
)

type User interface {
	Add(w http.ResponseWriter, r *http.Request)
}

type user struct {
	userService userService.User
	logger      logger.Logger
}

func New(userService userService.User, logger logger.Logger) User {
	return &user{userService: userService, logger: logger}
}
