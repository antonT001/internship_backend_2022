package transaction

import (
	"strings"
	"user_balance/service/internal/constants"
	"user_balance/service/internal/models"
)

func (u *transaction) List(input *models.TransactionListIn) ([]models.TransactionListFields, error) {
	transactionRows := []models.TransactionListFields{}
	str, whereArgs := getFilterList(input)
	err := u.db.Select(
		&transactionRows,
		`SELECT id, user_id, service_id, service_name, order_id, type, money, confirmed 
		FROM transactions WHERE `+str+` LIMIT ?, ?`,
		whereArgs...)

	if err != nil {
		return nil, err
	}

	return transactionRows, nil
}

func getFilterList(input *models.TransactionListIn) (string, []interface{}) {
	offset := input.PageNum.IntID() * constants.RESPONSE_LIMIT_DB
	whereArgs := make([]interface{}, 0)
	whereStrings := make([]string, 0)
	whereArgs = append(whereArgs, input.UserID.IntID())
	whereStrings = append(whereStrings, `user_id = ?`)

	if input.Filter == nil {
		whereArgs = append(whereArgs, offset)
		whereArgs = append(whereArgs, constants.RESPONSE_LIMIT_DB)
		str := strings.Join(whereStrings, " AND ")
		return str, whereArgs
	}

	//оставил возможномть добавить опции
	str := strings.Join(whereStrings, " AND ")

	if input.Filter.OrderBy != nil {
		str = str + " ORDER BY " + input.Filter.OrderBy.String()
	}

	if input.Filter.OrderDirection != nil {
		str = str + " " + input.Filter.OrderDirection.String()
	}

	whereArgs = append(whereArgs, offset)
	whereArgs = append(whereArgs, constants.RESPONSE_LIMIT_DB)

	return str, whereArgs
}
