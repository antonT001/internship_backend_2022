package balance

import (
	"database/sql"
	"fmt"
	"time"
	"user_balance/service/internal/models"
)

func (c *balance) Add(balance *models.TransactionFields) (sql.Result, error) {
	balance.CreatedAt = time.Now().Unix()
	fmt.Println("service")
	return c.hub.Balance().Add(balance)
}
