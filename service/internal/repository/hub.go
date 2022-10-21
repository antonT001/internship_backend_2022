package repository

import (
	"user_balance/service/internal/clients"
	"user_balance/service/internal/logger"
	"user_balance/service/internal/repository/accounting"
	"user_balance/service/internal/repository/balance"
	"user_balance/service/internal/repository/transaction"
)

type Hub interface {
	Accounting() accounting.Accounting
	Transaction() transaction.Transaction
	Balance() balance.Balance
}

type hub struct {
	accounting  accounting.Accounting
	transaction transaction.Transaction
	balance     balance.Balance
	logger      logger.Logger
}

func New(db clients.DataBase, logger logger.Logger) Hub {
	return &hub{
		accounting:  accounting.New(db),
		transaction: transaction.New(db),
		balance:     balance.New(db),
		logger:      logger,
	}
}

func (h *hub) Accounting() accounting.Accounting {
	return h.accounting
}

func (h *hub) Transaction() transaction.Transaction {
	return h.transaction
}

func (h *hub) Balance() balance.Balance {
	return h.balance
}
