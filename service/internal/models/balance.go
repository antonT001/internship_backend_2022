package models

type BalanceFields struct {
	User_ID uint64 `json:"user_id"`
	Money   uint64 `json:"money"`
}

type BalanceGetIn struct {
	UserID uint64 `json:"user_id"`
}

type BalanceGetOut struct {
	Success bool           `json:"success"`
	Balance *BalanceFields `json:"balance,omitempty"`
	Error   *string        `json:"error,omitempty"`
}
