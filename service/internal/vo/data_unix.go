package vo

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type DateUnix struct {
	value uint64
}

func ExamineDateUniix(value uint64) (DateUnix, error) {
	var (
		DateUnix DateUnix
		err      error
	)

	if value > uint64(time.Now().Unix()) || value < 1192406400 { //1192406400 - 15 Oct 2007 00:00:00 GMT
		err = fmt.Errorf("init timestamp error")
	}
	DateUnix.value = value

	return DateUnix, err
}

func (d DateUnix) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.value)
}

func (d *DateUnix) Scan(value interface{}) error {
	switch value := value.(type) {
	case uint64:
		d.value = value
		return nil
	default:
		return fmt.Errorf("cannot scan %T", value)
	}
}

func (d DateUnix) DateUnix() uint64 {
	return d.value
}

func (d DateUnix) Value() (driver.Value, error) {
	return d.value, nil
}
