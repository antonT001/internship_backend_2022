package transaction

import (
	"database/sql"
	"fmt"
	"user_balance/service/internal/models"
)

func (u *transaction) Confirm(input *models.TransactionConfirmFields) (result sql.Result, err error) {
	result, err = u.db.Exec("LOCK TABLES transactions WRITE")
	defer u.db.Exec("UNLOCK TABLES")
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

	err = checkType(tx, input)
	if err != nil {
		return nil, err
	}

	_, err = tx.NamedExec(`UPDATE transactions
	SET status=1, confirmed=:confirmed
	WHERE order_id=:order_id	
	`, *input)
	if err != nil {
		return nil, fmt.Errorf("failed to transaction confirm:%v", err)
	}

	result, err = tx.NamedExec(`INSERT INTO accounting (month, year, service_id, service_name, money)
	VALUES (:month, :year, :service_id, :service_name, :money) 
		ON DUPLICATE KEY 
	UPDATE money = money + VALUES(money)`, *input)
	if err != nil {
		return nil, fmt.Errorf("failed to add accounting data:%v", err)
	}

	return nil, tx.Commit()
}
