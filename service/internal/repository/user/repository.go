package user

import (
	"database/sql"
	"user_balance/service/internal/clients"
	"user_balance/service/internal/models"
)

type User interface {
	Add(product *models.UserFieldsAdd) (sql.Result, error)
}

type user struct {
	db clients.DataBase
}

func New(db clients.DataBase) User {
	return &user{
		db: db,
	}
}
