package accounting

import (
	"fmt"
	"user_balance/service/internal/models"
)

func (u *accounting) List(input *models.AccountingListIn) ([]models.AccountingListFields, error) {
	accountingRows := []models.AccountingListFields{}
	err := u.db.Select(&accountingRows,
		`SELECT service_name, money FROM accounting WHERE month = ? AND year = ?`,
		input.Month,
		input.Year,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get data for accounting: %v", err)
	}
	return accountingRows, nil
}
