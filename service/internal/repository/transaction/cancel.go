package transaction

import (
	"database/sql"
	"fmt"
	"user_balance/service/internal/models"
)

func (u *transaction) Cancel(input *models.TransactionConfirmFields) (result sql.Result, err error) {
	u.db.Exec("LOCK TABLES balance WRITE, transactions WRITE")
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

	result, err = tx.NamedExec(`UPDATE transactions
	SET status=2, confirmed=:confirmed
	WHERE order_id=:order_id	
	`, *input)
	if err != nil {
		return nil, fmt.Errorf("failed to transaction cancel:%v", err)
	}

	result, err = tx.NamedExec(`UPDATE balance
	SET money = money + :money
	WHERE user_id = :user_id	
	`, input)
	if err != nil {
		return nil, fmt.Errorf("failed to return balance:%v", err)
	}

	return nil, tx.Commit()
}
