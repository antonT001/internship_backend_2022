package transaction

import (
	"fmt"
	"user_balance/service/internal/models"
)

func (c *transaction) List(input *models.TransactionListIn) ([]models.TransactionListFields, error) {
	fmt.Println("service")
	return c.hub.Transaction().List(input)
}
