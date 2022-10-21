package balance

import (
	"fmt"
	"user_balance/service/internal/models"
	"user_balance/service/internal/vo"
)

func (u *balance) Get(userId *vo.IntID) (*models.BalanceFields, error) {
	balance := models.BalanceFields{}
	err := u.db.Get(
		&balance,
		`SELECT user_id, money FROM balance WHERE user_id = ?`,
		userId)

	if err != nil {
		return nil, fmt.Errorf("unable to get balance: %v", err)
	}
	return &balance, nil
}
