package models

import "user_balance/service/internal/vo"

type UserFieldsAdd struct {
	UserName vo.Name  `db:"user_name"`
}
type UserAdd struct {
	UserName string `json:"user_name"`
}
