package vo

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type DeltaMoney struct {
	value float64
}

func ExamineDeltaMoney(value float64) (DeltaMoney, error) {
	var (
		id  DeltaMoney
		err error
	)

	id.value = value

	if id.value == 0.0 {
		err = fmt.Errorf("delta money error")
	}

	return id, err
}

func (id DeltaMoney) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.value)
}

func (id *DeltaMoney) Scan(value interface{}) error {
	switch value := value.(type) {
	case float64:
		id.value = value
		return nil
	default:
		return fmt.Errorf("cannot scan %T", value)
	}
}

func (id DeltaMoney) DeltaMoney() float64 {
	return id.value
}

func (id DeltaMoney) Value() (driver.Value, error) {
	return id.value, nil
}
