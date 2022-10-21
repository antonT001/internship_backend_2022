package transaction

import (
	"database/sql"
	"user_balance/service/internal/clients"
	"user_balance/service/internal/models"
)

type Transaction interface {
	Pay(input *models.TransactionFields) (result sql.Result, err error)
	Confirm(input *models.TransactionConfirmFields) (sql.Result, error)
	Cancel(input *models.TransactionConfirmFields) (result sql.Result, err error)
	List(input *models.TransactionListIn) ([]models.TransactionListFields, error)
}

type transaction struct {
	db clients.DataBase
}

func New(db clients.DataBase) Transaction {
	return &transaction{
		db: db,
	}
}
