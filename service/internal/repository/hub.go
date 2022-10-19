package repository

import (
	"user_balance/service/internal/clients"
	"user_balance/service/internal/logger"
	"user_balance/service/internal/repository/balance"
	"user_balance/service/internal/repository/user"
)

type Hub interface {
	User() user.User
	Balance() balance.Balance
}

type hub struct {
	user    user.User
	balance balance.Balance
	logger  logger.Logger
}

func New(db clients.DataBase, logger logger.Logger) Hub {
	return &hub{
		user:   user.New(db),
		balance: balance.New(db),
		logger: logger,
	}
}

func (h *hub) User() user.User {
	return h.user
}

func (h *hub) Balance() balance.Balance {
	return h.balance
}
