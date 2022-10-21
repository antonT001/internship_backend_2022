package accounting

import (
	"user_balance/service/internal/logger"
	"user_balance/service/internal/models"
	"user_balance/service/internal/repository"
)

type Accounting interface {
	List(input *models.AccountingListIn) ([]models.AccountingListFields, error)
}

type accounting struct {
	hub    repository.Hub
	logger logger.Logger
}

func New(hub repository.Hub, logger logger.Logger) Accounting {
	return &accounting{hub: hub, logger: logger}
}
