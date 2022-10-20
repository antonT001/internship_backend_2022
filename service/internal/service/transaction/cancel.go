package transaction

import (
	"database/sql"
	"fmt"
	"time"
	"user_balance/service/internal/models"
)

func (c *transaction) Cancel(transaction *models.TransactionConfirmFields) (sql.Result, error) {
	transaction.UpdateAT = time.Now().Unix()
	fmt.Println("service")
	return c.hub.Transaction().Cancel(transaction)
}
