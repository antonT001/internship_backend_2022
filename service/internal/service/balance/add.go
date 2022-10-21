package balance

import (
	"database/sql"
	"fmt"
	"time"
	"user_balance/service/internal/models"
)

func (c *balance) Add(input *models.TransactionFields) (sql.Result, error) {
	input.CreatedAt = time.Now().Unix()
	input.Confirmed = input.CreatedAt
	fmt.Println("service")
	return c.hub.Balance().Add(input)
}
