package vo

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type OrderDirection struct {
	value *string
}

func ExamineOrderDirection(value string) (*OrderDirection, error) {
	var (
		OrderDirection OrderDirection
		err            error
	)

	if value == "ASC" || value == "DESC" {
		OrderDirection.value = &value
		return &OrderDirection, nil
	}

	err = fmt.Errorf(" error sort_direction")
	return &OrderDirection, err
}

func ExaminePointerOrderDirection(value *string) (*OrderDirection, error) {
	if value == nil {
		return nil, nil
	}

	return ExamineOrderDirection(*value)
}

func (id OrderDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.value)
}

func (id *OrderDirection) Scan(value interface{}) error {
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

func (id OrderDirection) String() string {
	return *id.value
}

func (id OrderDirection) Value() (driver.Value, error) {
	if id.value == nil {
		return nil, nil
	}

	return *id.value, nil
}
