package balance

import (
	"database/sql"
	"fmt"
	"user_balance/service/internal/models"
)

func (u *balance) Add(balance *models.TransactionFields) (result sql.Result, err error) {
	tx, _ := u.db.NewTransaction()
	defer func() {
		if err == nil {
			return
		}
		errs := tx.Rollback()
		if errs != nil {
			err = errs
		}
	}()

	result, err = tx.NamedExec(`INSERT INTO balance (user_id, money)
	VALUES (:user_id, :money) 
		ON DUPLICATE KEY 
	UPDATE money = money + VALUES(money)`, *balance)
	if err != nil {
		return nil, fmt.Errorf("failed to add balance:%v", err)
	}

	result, err = tx.NamedExec(`INSERT INTO transactions 
	(user_id, service_id, service_name, order_id, type, money, status, created_at)
	VALUES (:user_id, :service_id, :service_name, :order_id, :type, :money, 1, :created_at)`,
		*balance)
	if err != nil {
		return nil, fmt.Errorf("failed to add transaction:%v", err)
	}

	return nil, tx.Commit()
}
