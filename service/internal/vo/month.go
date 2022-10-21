package vo

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Month struct {
	value uint64
}

func ExamineMonth(value uint64) (Month, error) {
	var (
		Month Month
		err   error
	)

	if value > 12 || value < 1 {
		err = fmt.Errorf("month error")
	}
	Month.value = value

	return Month, err
}

func (d Month) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.value)
}

func (d *Month) Scan(value interface{}) error {
	switch value := value.(type) {
	case uint64:
		d.value = value
		return nil
	default:
		return fmt.Errorf("cannot scan %T", value)
	}
}

func (d Month) Month() uint64 {
	return d.value
}

func (d Month) Value() (driver.Value, error) {
	return d.value, nil
}
