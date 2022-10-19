package balance

import (
	"database/sql"
	"user_balance/service/internal/clients"
	"user_balance/service/internal/models"
)

type Balance interface {
	Add(balance *models.BalanceFields) (sql.Result, error)
	Pay(balance *models.BalanceFields) (result sql.Result, err error)
}

type balance struct {
	db clients.DataBase
}

func New(db clients.DataBase) Balance {
	return &balance{
		db: db,
	}
}
