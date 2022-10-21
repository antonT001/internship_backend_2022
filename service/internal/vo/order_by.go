package vo

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"user_balance/service/internal/helpers"
)

var (
	OrderByTransactionList = helpers.MakeMapEmptyValue(
		[]string{"confirmed", "money"},
	)
)

type OrderBy struct {
	value *string
}

func ExamineOrderBy(value string, m map[string]struct{}) (*OrderBy, error) {
	var (
		OrderBy OrderBy
		err     error
	)

	if _, ok := m[value]; ok {
		OrderBy.value = &value
		return &OrderBy, nil

	}
	err = fmt.Errorf(" error order_by")
	return &OrderBy, err
}

func ExaminePointerOrderBy(value *string, m map[string]struct{}) (*OrderBy, error) {
	if value == nil {
		return nil, nil
	}

	return ExamineOrderBy(*value, m)
}

func (id OrderBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.value)
}

func (id *OrderBy) Scan(value interface{}) error {
	switch value := value.(type) {
	case []uint8:
		a := string(value)
		id.value = &a
		return nil
	case nil:
		return nil
	default:
		return fmt.Errorf("cannot scan %T", value)
	}
}

func (id OrderBy) String() string {
	return *id.value
}

func (id OrderBy) Value() (driver.Value, error) {
	if id.value == nil {
		return nil, nil
	}

	return *id.value, nil
}
