package core

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type EntityDate time.Time

func (date *EntityDate) Scan(value any) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*date = EntityDate(nullTime.Time)
	return
}

func (date EntityDate) Value() (driver.Value, error) {
	t := time.Time(date)
	if t.IsZero() {
		return nil, nil
	}
	y, m, d := time.Time(date).Date()
	return time.Date(y, m, d, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Time(date).Location()), nil
}

// GormDataType gorm common data type
func (date EntityDate) GormDataType() string {
	return "timestamp"
}

func (date EntityDate) GobEncode() ([]byte, error) {
	return time.Time(date).GobEncode()
}

func (date *EntityDate) GobDecode(b []byte) error {
	return (*time.Time)(date).GobDecode(b)
}

func (date EntityDate) MarshalJSON() ([]byte, error) {
	t := time.Time(date)
	if t.IsZero() {
		return []byte("0"), nil
	}
	return []byte(fmt.Sprintf("%d", t.UnixMicro()/1e3)), nil
}

func (date *EntityDate) UnmarshalJSON(b []byte) error {
	return (*time.Time)(date).UnmarshalJSON(b)
}

func (date EntityDate) GetTime() time.Time {
	t := time.Time(date)
	return t
}
