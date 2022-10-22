package accounting

import (
	"user_balance/service/internal/logger"
	"user_balance/service/internal/models"
	"user_balance/service/internal/objectstorage"
	"user_balance/service/internal/repository"
)

type Accounting interface {
	List(input *models.AccountingListIn) (*string, error)
}

type accounting struct {
	hub           repository.Hub
	objectStorage objectstorage.ObjectStorage
	logger        logger.Logger
}

func New(hub repository.Hub, objectStorage objectstorage.ObjectStorage, logger logger.Logger) Accounting {
	return &accounting{hub: hub, objectStorage: objectStorage, logger: logger}
}
