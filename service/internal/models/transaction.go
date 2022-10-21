package models

import "user_balance/service/internal/vo"

type TransactionFields struct {
	UserID      vo.IntID      `db:"user_id"`
	Type        int           `db:"type"` //TODO добавить vo, Type: 0 - списание, 1 - пополнение
	Money       vo.DeltaMoney `db:"money"`
	ServiceID   vo.IntID      `db:"service_id"`
	ServiceName vo.Name       `db:"service_name"`
	OrderID     vo.IntID      `db:"order_id"`
	Confirmed   int64         `db:"confirmed"`
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
	UserID      vo.IntID      `db:"user_id"`
	ServiceID   vo.IntID      `db:"service_id"`
	ServiceName vo.Name       `db:"service_name"`
	OrderID     vo.IntID      `db:"order_id"`
	Money       vo.DeltaMoney `db:"money"`
	Confirmed   int64         `db:"confirmed"`
	Month       int           `db:"month"`
	Year        int           `db:"year"`
}

type TransactionConfirm struct {
	UserID      uint64 `json:"user_id"`
	ServiceID   uint64 `json:"service_id"`
	ServiceName string `json:"service_name"`
	OrderID     uint64 `json:"order_id"`
	Money       uint64 `json:"money"`
}

type TransactionListFields struct {
	ID           uint64 `json:"id"`
	User_ID      uint64 `json:"user_id"`
	Service_ID   uint64 `json:"service_id"`
	Service_Name string `json:"service_name"`
	Order_ID     uint64 `json:"order_id"`
	Type         int    `json:"type"`
	Money        uint64 `json:"money"`
	Confirmed    int64  `json:"confirmed"`
}

type TransactionFilter struct {
	OrderBy        string `json:"order_by"`
	OrderDirection string `json:"order_direction"`
}

type TransactionList struct {
	PageNum *uint64            `json:"page_num"`
	UserID  uint64             `json:"user_id"`
	Filter  *TransactionFilter `json:"filter"`
}

type TransactionFilterIn struct {
	OrderBy        *vo.OrderBy
	OrderDirection *vo.OrderDirection
}

type TransactionListIn struct {
	PageNum vo.IntID
	UserID  vo.IntID
	Filter  *TransactionFilterIn
}

type TransactionListOut struct {
	Success     bool                    `json:"success"`
	Transaction []TransactionListFields `json:"transaction_list,omitempty"`
	Error       *string                 `json:"error,omitempty"`
}
