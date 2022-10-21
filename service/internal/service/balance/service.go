package balance

import (
	"database/sql"
	"user_balance/service/internal/logger"
	"user_balance/service/internal/models"
	"user_balance/service/internal/repository"
	"user_balance/service/internal/vo"
)

type Balance interface {
	Add(input *models.TransactionFields) (sql.Result, error)
	Get(userId *vo.IntID) (*models.BalanceFields, error)
}

type balance struct {
	hub    repository.Hub
	logger logger.Logger
}

func New(hub repository.Hub, logger logger.Logger) Balance {
	return &balance{hub: hub, logger: logger}
}
