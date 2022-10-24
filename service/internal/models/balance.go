package models

type BalanceFields struct {
	User_ID uint64 `json:"user_id" example:"123"`
	// Money:
	// * The format of money without kopecks is used.
	// * Example 12050=120 rubles 50 kopecks
	Money uint64 `json:"money"`
}

type BalanceGetIn struct {
	UserID uint64 `json:"user_id" example:"123"`
}

type BalanceGetOut struct {
	Success bool           `json:"success"`
	Balance *BalanceFields `json:"balance,omitempty"`
	Error   *string        `json:"error,omitempty"`
}
