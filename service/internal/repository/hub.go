package repository

import (
	"user_balance/service/internal/clients"
	"user_balance/service/internal/logger"
	"user_balance/service/internal/repository/user"
)

type Hub interface {
	User() user.User
}

type hub struct {
	user   user.User
	logger logger.Logger
}

func New(db clients.DataBase, logger logger.Logger) Hub {
	return &hub{
		user:   user.New(db),
		logger: logger,
	}
}

func (h *hub) User() user.User {
	return h.user
}
