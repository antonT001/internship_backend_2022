package balance

import (
	"database/sql"
	"user_balance/service/internal/logger"
	"user_balance/service/internal/models"
	"user_balance/service/internal/repository"
)

type Balance interface {
	Add(balance *models.BalanceFieldsAdd) (sql.Result, error)
}

type balance struct {
	hub    repository.Hub
	logger logger.Logger
}

func New(hub repository.Hub, logger logger.Logger) Balance {
	return &balance{hub: hub, logger: logger}
}
