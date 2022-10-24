package transaction

//go:generate ../../../../bin/mockery --name=Transaction --case underscore

import (
	"database/sql"
	"user_balance/service/internal/logger"
	"user_balance/service/internal/models"
	"user_balance/service/internal/repository"
)

type Transaction interface {
	Add(input *models.TransactionFields) (sql.Result, error)
	Confirm(input *models.TransactionConfirmFields) (sql.Result, error)
	Cancel(input *models.TransactionConfirmFields) (sql.Result, error)
	List(input *models.TransactionListIn) ([]models.TransactionListFields, error)
}

type transaction struct {
	hub    repository.Hub
	logger logger.Logger
}

func New(hub repository.Hub, logger logger.Logger) Transaction {
	return &transaction{hub: hub, logger: logger}
}
