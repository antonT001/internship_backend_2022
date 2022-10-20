package transaction

import (
	"database/sql"
	"user_balance/service/internal/clients"
	"user_balance/service/internal/models"
)

type Transaction interface {
	Pay(transaction *models.TransactionFields) (result sql.Result, err error)
	Confirm(transaction *models.TransactionConfirmFields) (sql.Result, error)
	Cancel(transaction *models.TransactionConfirmFields) (result sql.Result, err error)
}

type transaction struct {
	db clients.DataBase
}

func New(db clients.DataBase) Transaction {
	return &transaction{
		db: db,
	}
}
