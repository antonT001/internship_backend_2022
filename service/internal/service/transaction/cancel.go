package transaction

import (
	"database/sql"
	"fmt"
	"time"
	"user_balance/service/internal/models"
)

func (c *transaction) Cancel(input *models.TransactionConfirmFields) (sql.Result, error) {
	input.Confirmed = time.Now().Unix()
	fmt.Println("service")
	return c.hub.Transaction().Cancel(input)
}
