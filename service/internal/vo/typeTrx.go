package vo

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type TypeTrx struct {
	value uint64
}

func ExamineTypeTrx(value uint64) (TypeTrx, error) {
	var (
		id  TypeTrx
		err error
	)

	id.value = value

	if id.value > 1 {
		err = fmt.Errorf("type transaction error")
	}

	return id, err
}

func (id TypeTrx) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.value)
}

func (id *TypeTrx) Scan(value interface{}) error {
	switch value := value.(type) {
	case uint64:
		id.value = value
		return nil
	default:
		return fmt.Errorf("cannot scan %T", value)
	}
}

func (id TypeTrx) TypeTrx() uint64 {
	return id.value
}

func (id TypeTrx) Value() (driver.Value, error) {
	return id.value, nil
}
