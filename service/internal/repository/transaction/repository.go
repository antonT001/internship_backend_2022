package transaction

import (
	"database/sql"
	"fmt"
	"user_balance/service/internal/clients"
	"user_balance/service/internal/constants"
	"user_balance/service/internal/models"
)

type Transaction interface {
	Pay(input *models.TransactionFields) (result sql.Result, err error)
	Confirm(input *models.TransactionConfirmFields) (sql.Result, error)
	Cancel(input *models.TransactionConfirmFields) (result sql.Result, err error)
	List(input *models.TransactionListIn) ([]models.TransactionListFields, error)
}

type transaction struct {
	db clients.DataBase
}

func New(db clients.DataBase) Transaction {
	return &transaction{
		db: db,
	}
}

func checkType(tx clients.Transaction, input *models.TransactionConfirmFields) error {
	var status int
	err := tx.Get(
		&status,
		`SELECT status FROM transactions 
		WHERE user_id = ? AND service_id = ? AND order_id = ? AND money = ?`,
		input.UserID,
		input.ServiceID,
		input.OrderID,
		input.Money,
	)
	if err != nil {
		return fmt.Errorf("failed to get transaction information:%v", err)
	}

	switch status {
	case constants.STATUS_CANCEL:
		return fmt.Errorf("transaction has already been canceled")
	case constants.STATUS_CONFIRM:
		return fmt.Errorf("transaction has already been confirmed")
	case constants.STATUS_RESERVED:
		return nil
	default:
		return fmt.Errorf("unknown transaction status - %v", status)
	}
}
