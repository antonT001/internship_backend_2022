package models

import "user_balance/service/internal/vo"

type UserFieldsAdd struct {
	ID      vo.IntID      `db:"id"`
	Balance vo.DeltaMoney `db:"balance"`
}
type UserAdd struct {
	ID      uint64  `json:"id"`
	Balance float64 `json:"balance"`
}
