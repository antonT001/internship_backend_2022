package transaction

import (
	"database/sql"
	"fmt"
	"user_balance/service/internal/constants"
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

	var status int //TODO - повторяющийся кусок кода с service/internal/repository/balance/cancel.go
	err = tx.Get(
		&status,
		`SELECT status FROM transactions 
		WHERE user_id = ? AND service_id = ? AND order_id = ? AND money = ?`,
		input.UserID,
		input.ServiceID,
		input.OrderID,
		input.Money,
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

	_, err = tx.NamedExec(`UPDATE transactions
	SET status=1, confirmed=:confirmed
	WHERE order_id=:order_id	
	`, *input)
	if err != nil {
		return nil, fmt.Errorf("failed to transaction confirm:%v", err)
	}

	return nil, tx.Commit()
}
