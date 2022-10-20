package transaction

import (
	"database/sql"
	"fmt"
	"time"
	"user_balance/service/internal/models"
)

func (c *transaction) Confirm(transaction *models.TransactionConfirmFields) (sql.Result, error) {
	transaction.UpdateAT = time.Now().Unix()
	fmt.Println("service")
	return c.hub.Transaction().Confirm(transaction)
}
