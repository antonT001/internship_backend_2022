package repository

import (
	"user_balance/service/internal/clients"
	"user_balance/service/internal/logger"
	"user_balance/service/internal/repository/balance"
	"user_balance/service/internal/repository/transaction"
	"user_balance/service/internal/repository/user"
)

type Hub interface {
	User() user.User
	Transaction() transaction.Transaction
	Balance() balance.Balance
}

type hub struct {
	user        user.User
	transaction transaction.Transaction
	balance     balance.Balance
	logger      logger.Logger
}

func New(db clients.DataBase, logger logger.Logger) Hub {
	return &hub{
		user:        user.New(db),
		transaction: transaction.New(db),
		balance:     balance.New(db),
		logger:      logger,
	}
}

func (h *hub) User() user.User {
	return h.user
}

func (h *hub) Transaction() transaction.Transaction {
	return h.transaction
}

func (h *hub) Balance() balance.Balance {
	return h.balance
}
