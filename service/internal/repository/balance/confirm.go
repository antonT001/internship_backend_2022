package balance

import (
	"database/sql"
	"fmt"
	"user_balance/service/internal/models"
)

func (u *balance) Confirm(balance *models.BalanceConfirmFields) (sql.Result, error) {
	_, err := u.db.NamedExec(`UPDATE transactions
	SET confirm=1
	WHERE process_id = :process_id	
	`, *balance)
	if err != nil {
		return nil, fmt.Errorf("failed to transactions confirm:%v", err)
	}

	return nil, nil
}
