package transaction

import (
	"database/sql"
	"user_balance/service/internal/logger"
	"user_balance/service/internal/models"
	"user_balance/service/internal/repository"
)

type Transaction interface {
	Pay(transaction *models.TransactionFields) (sql.Result, error)
	Confirm(transaction *models.TransactionConfirmFields) (sql.Result, error)
	Cancel(transaction *models.TransactionConfirmFields) (sql.Result, error)
}

type transaction struct {
	hub    repository.Hub
	logger logger.Logger
}

func New(hub repository.Hub, logger logger.Logger) Transaction {
	return &transaction{hub: hub, logger: logger}
}