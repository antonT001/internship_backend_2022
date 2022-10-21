package models

import "user_balance/service/internal/vo"

type UserFieldsAdd struct {
	UserName vo.Name `db:"user_name"`
}
type UserAdd struct {
	UserName string `json:"user_name"`
}

type TransactionFields struct {
	UserID      vo.IntID      `db:"user_id"`
	Type        int           `db:"type"` //TODO добавить vo, Type: 0 - списание, 1 - пополнение
	Money       vo.DeltaMoney `db:"money"`
	ServiceID   vo.IntID      `db:"service_id"`
	ServiceName vo.Name       `db:"service_name"`
	OrderID     vo.IntID      `db:"order_id"`
	CreatedAt   int64         `db:"created_at"`
}
type Transaction struct {
	UserID      uint64 `json:"user_id"`
	Type        int    `json:"type"`
	Money       uint64 `json:"money"`
	ServiceID   uint64 `json:"service_id"`
	ServiceName string `json:"service_name"`
	OrderID     uint64 `json:"order_id"`
}

type TransactionConfirmFields struct {
	UserID    vo.IntID      `db:"user_id"`
	ServiceID vo.IntID      `db:"service_id"`
	OrderID   vo.IntID      `db:"order_id"`
	Money     vo.DeltaMoney `db:"money"`
	UpdateAT  int64         `db:"update_at"`
}

type TransactionConfirm struct {
	UserID    uint64 `json:"user_id"`
	ServiceID uint64 `json:"service_id"`
	OrderID   uint64 `json:"order_id"`
	Money     uint64 `json:"money"`
}

type BalanceFields struct {
	UserID uint64 `db:"user_id"`
	Money  uint64 `db:"money"`
}

type BalanceGetIn struct {
	UserID uint64 `json:"user_id"`
}

type BalanceGetOut struct {
	Success bool           `json:"success"`
	Balance *BalanceFields `json:"balance,omitempty"`
	Error   *string        `json:"error,omitempty"`
}
