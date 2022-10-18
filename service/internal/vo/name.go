package vo

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Name struct {
	value *string
}

func ExamineName(value string) (*Name, error) {
	var (
		Name Name
		err  error
	)

	if value == "" {
		err = fmt.Errorf("name error")
	}

	Name.value = &value

	return &Name, err
}

func ExaminePointerName(value *string) (*Name, error) {
	if value == nil {
		return nil, nil
	}
	return ExamineName(*value)
}

func (id Name) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.value)
}

func (id *Name) Scan(value interface{}) error {
	switch value := value.(type) {
	case []uint8:
		a := string(value)
		id.value = &a
		return nil
	default:
		return fmt.Errorf("cannot scan %T", value)
	}
}

func (id Name) String() string {
	return *id.value
}

func (id Name) Value() (driver.Value, error) {
	return *id.value, nil
}
