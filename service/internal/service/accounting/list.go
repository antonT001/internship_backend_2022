package accounting

import (
	"user_balance/service/internal/models"
)

func (c *accounting) List(input *models.AccountingListIn) ([]models.AccountingListFields, error) {
	return c.hub.Accounting().List(input)
}
