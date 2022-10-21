package transaction

import (
	"database/sql"
	"fmt"
	"time"
	"user_balance/service/internal/models"
)

func (c *transaction) Pay(input *models.TransactionFields) (sql.Result, error) {
	input.CreatedAt = time.Now().Unix()
	fmt.Println("service")
	return c.hub.Transaction().Pay(input)
}
