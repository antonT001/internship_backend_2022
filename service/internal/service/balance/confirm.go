package balance

import (
	"database/sql"
	"fmt"
	"user_balance/service/internal/models"
)

func (c *balance) Confirm(balance *models.BalanceConfirmFields) (sql.Result, error) {
	fmt.Println("service")
	return c.hub.Balance().Confirm(balance)
}
