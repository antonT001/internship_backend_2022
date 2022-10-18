package vo

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type IntID struct {
	value uint64
}

func ExamineIntID(value uint64) (IntID, error) {
	var (
		id  IntID
		err error
	)

	id.value = value

	if id.value < 1 {
		err = fmt.Errorf("id error")
	}

	return id, err
}

func ExaminePointerIntID(value *uint64) (IntID, error) {
	if value == nil {
		return IntID{0}, nil
	}

	return ExamineIntID(*value)
}

func (id IntID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.value)
}

func (id *IntID) Scan(value interface{}) error {
	switch value := value.(type) {
	case uint64:
		id.value = value
		return nil
	default:
		return fmt.Errorf("cannot scan %T", value)
	}
}

func (id IntID) IntID() uint64 {
	return id.value
}

func (id IntID) Value() (driver.Value, error) {
	return id.value, nil
}
