package transaction

import (
	"database/sql"
	"fmt"
	"user_balance/service/internal/constants"
	"user_balance/service/internal/models"
)

func (u *transaction) Cancel(transaction *models.TransactionConfirmFields) (result sql.Result, err error) {
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

	var status int //TODO - повторяющийся кусок кода с service/internal/repository/balance/cancel.go
	err = tx.Get(
		&status,
		`SELECT status FROM transactions 
		WHERE user_id = ? AND service_id = ? AND order_id = ? AND money = ?`,
		transaction.UserID,
		transaction.ServiceID,
		transaction.OrderID,
		transaction.Money,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction information:%v", err)
	}
	fmt.Printf("status: %v\n", status)
	switch status {
	case constants.STATUS_CANCEL:
		return nil, fmt.Errorf("transaction has already been canceled")
	case constants.STATUS_CONFIRM:
		return nil, fmt.Errorf("transaction has already been confirmed")
	case constants.STATUS_RESERVED:
		/*...*/
	default:
		return nil, fmt.Errorf("unknown transaction status - %v", status)
	}
	result, err = tx.NamedExec(`UPDATE transactions
	SET status=2, update_at=:update_at
	WHERE order_id=:order_id	
	`, *transaction)
	if err != nil {
		return nil, fmt.Errorf("failed to transaction cancel:%v", err)
	}

	result, err = tx.NamedExec(`UPDATE balance
	SET money = money + :money
	WHERE user_id = :user_id	
	`, transaction)
	if err != nil {
		return nil, fmt.Errorf("failed to return balance:%v", err)
	}

	return nil, tx.Commit()
}