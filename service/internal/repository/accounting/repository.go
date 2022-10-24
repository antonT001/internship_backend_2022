package accounting

//go:generate ../../../../bin/mockery --name=Accounting --case underscore

import (
	"user_balance/service/internal/clients"
	"user_balance/service/internal/models"
)

type Accounting interface {
	List(input *models.AccountingListIn) ([]models.AccountingListFields, error)
}

type accounting struct {
	db clients.DataBase
}

func New(db clients.DataBase) Accounting {
	return &accounting{
		db: db,
	}
}
