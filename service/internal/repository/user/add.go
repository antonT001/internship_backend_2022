package user

import (
	"database/sql"
	"user_balance/service/internal/models"
)

func (c *user) Add(user *models.UserFieldsAdd) (sql.Result, error) {
	return nil, nil
}
