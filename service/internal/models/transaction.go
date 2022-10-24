package models

import "user_balance/service/internal/vo"

type TransactionFields struct {
	UserID      vo.IntID      `db:"user_id"`
	Type        vo.TypeTrx    `db:"type"`
	Money       vo.DeltaMoney `db:"money"`
	ServiceID   vo.IntID      `db:"service_id"`
	ServiceName vo.Name       `db:"service_name"`
	OrderID     vo.IntID      `db:"order_id"`
	Confirmed   int64         `db:"confirmed"`
	CreatedAt   int64         `db:"created_at"`
}

type Transaction struct {
	UserID uint64 `json:"user_id" example:"123"`
	// Type transaction:
	// * 0 - withdrawal of money from the user's account
	// * 1 - receipt of money on the user's account
	Type uint64 `json:"type" enums:"0,1"`
	// Money:
	// * The format of money without kopecks is used.
	// * Example 12050=120 rubles 50 kopecks
	Money       uint64 `json:"money"`
	ServiceID   uint64 `json:"service_id" example:"123"`
	ServiceName string `json:"service_name"`
	OrderID     uint64 `json:"order_id" example:"123"`
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
	UserID      uint64 `json:"user_id" example:"123"`
	ServiceID   uint64 `json:"service_id" example:"123"`
	ServiceName string `json:"service_name"`
	OrderID     uint64 `json:"order_id" example:"123"`
	// Money:
	// * The format of money without kopecks is used.
	// * Example 12050=120 rubles 50 kopecks
	Money uint64 `json:"money"`
}

type TransactionListFields struct {
	ID           uint64 `json:"id"`
	User_ID      uint64 `json:"user_id"`
	Service_ID   uint64 `json:"service_id"`
	Service_Name string `json:"service_name"`
	Order_ID     uint64 `json:"order_id"`
	// Type transaction:
	// * 0 - withdrawal of money from the user's account
	// * 1 - receipt of money on the user's account
	Type  uint64 `json:"type" enums:"0,1"`
	// Money:
	// * The format of money without kopecks is used.
	// * Example 12050=120 rubles 50 kopecks
	Money     uint64 `json:"money"`
	Confirmed int64  `json:"confirmed"`
}

type TransactionFilter struct {
	// Sortable field
	// * confirmed - data confirmed transaction
	// * money - transaction money amount
	OrderBy string `json:"order_by" enums:"confirmed,money"`
	// Sort order:
	// * ASC - Ascending, from A to Z.
	// * DESC - Descending, from Z to A.
	OrderDirection string `json:"order_direction" enums:"ASC,DESC"`
}

type TransactionList struct {
	// Pagination:
	// * default page number = 0, the first n lines from the list are displayed.
	// * If page number =1 then the following n lines are displayed
	// * where n is a constant declared in the constants package - RESPONSE_LIMIT_DB
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
