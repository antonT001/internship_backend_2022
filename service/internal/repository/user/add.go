package user

import (
	"database/sql"
	"fmt"
	"user_balance/service/internal/models"
)

func (u *user) Add(user *models.UserFieldsAdd) (sql.Result, error) {
	_, err := u.db.NamedExec(`INSERT INTO user (user_name) VALUES (:user_name)`, *user)
	if err != nil {
		return nil, fmt.Errorf("failed to add user:%v", err)
	}
	return nil, nil
}