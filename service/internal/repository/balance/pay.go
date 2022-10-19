package balance

import (
	"database/sql"
	"fmt"
	"user_balance/service/internal/models"
)

func (u *balance) Pay(balance *models.BalanceFields) (result sql.Result, err error) {
	u.db.Exec("LOCK TABLES balance READ, transaction READ")
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

	//получаем актуальный баланс

	summHold := 0
	err = tx.Get(
		&summHold,
		`SELECT money FROM balance WHERE user_id = ?`,
		balance.UserID)

	if err != nil {
		return nil, err
	}
	fmt.Printf("summ: %v\n", summHold)
	//сравниваем с суммой списания
	if balance.Money.DeltaMoney() > uint64(summHold) {
		return nil, fmt.Errorf("not enough money on balance")
	}
	//создаем заморозку, уменьшеаем баланс на эту сумму или возвращаем ошибру

	result, err = tx.NamedExec(`INSERT INTO transactions 
	(user_id, service_id, service_name, process_id, type, money, created_at)
	VALUES (:user_id, :service_id, :service_name, :process_id, :type, :money, :created_at)`,
		*balance)
	if err != nil {
		return nil, fmt.Errorf("failed to add transaction:%v", err)
	}

	result, err = tx.NamedExec(`UPDATE balance
	SET money= money - :money
	WHERE user_id = :user_id	
	`, *balance)
	if err != nil {
		return nil, fmt.Errorf("failed to add balance:%v", err)
	}

	return nil, tx.Commit()
}
