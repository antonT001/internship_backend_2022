package transaction

import (
	"database/sql"
	"fmt"
	"time"
	"user_balance/service/internal/models"
)

func (c *transaction) Pay(transaction *models.TransactionFields) (sql.Result, error) {
	transaction.CreatedAt = time.Now().Unix()
	fmt.Println("service")
	return c.hub.Transaction().Pay(transaction)
}
