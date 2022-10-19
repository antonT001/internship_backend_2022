package balance

import (
	"database/sql"
	"fmt"
	"time"
	"user_balance/service/internal/models"
)

func (c *balance) Pay(balance *models.BalanceFields) (sql.Result, error) {
	balance.CreatedAt = time.Now().Unix()
	fmt.Println("service")
	return c.hub.Balance().Pay(balance)
}
