package balance

import (
	"fmt"
	"user_balance/service/internal/models"
	"user_balance/service/internal/vo"
)

func (c *balance) Get(userId *vo.IntID) (*models.BalanceFields, error) {
	fmt.Println("service")
	return c.hub.Balance().Get(userId)
}
