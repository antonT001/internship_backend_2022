package vo

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type Year struct {
	value uint64
}

func ExamineYear(value uint64) (Year, error) {
	var (
		Year Year
		err  error
	)

	if value > uint64(time.Now().Year()) || value < 2017 {
		err = fmt.Errorf("year error")
	}
	Year.value = value

	return Year, err
}

func (d Year) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.value)
}

func (d *Year) Scan(value interface{}) error {
	switch value := value.(type) {
	case uint64:
		d.value = value
		return nil
	default:
		return fmt.Errorf("cannot scan %T", value)
	}
}

func (d Year) Year() uint64 {
	return d.value
}

func (d Year) Value() (driver.Value, error) {
	return d.value, nil
}
