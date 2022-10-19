package models

import "user_balance/service/internal/vo"

type UserFieldsAdd struct {
	UserName vo.Name `db:"user_name"`
}
type UserAdd struct {
	UserName string `json:"user_name"`
}

type BalanceFieldsAdd struct {
	UserID vo.IntID `db:"user_id"`
	//Type: 0 - списание, 1 - пополнение
	Type        int           `db:"type"` //TODO добавить vo
	Money       vo.DeltaMoney `db:"money"`
	ServiceID   vo.IntID      `db:"service_id"`
	ServiceName vo.Name       `db:"service_name"`
	ProcessID   vo.IntID      `db:"process_id"`
	CreatedAt   int64   `db:"created_at"`
}
type BalanceAdd struct {
	UserID uint64 `json:"user_id"`
	//Type: 0 - списание, 1 - пополнение
	Type        int    `json:"type"`
	Money       uint64 `json:"money"`
	ServiceID   uint64 `json:"service_id"`
	ServiceName string `json:"service_name"`
	ProcessID   uint64 `json:"process_id"`
}
