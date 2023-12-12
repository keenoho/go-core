package core

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type EntityDate time.Time

func (date *EntityDate) Scan(value any) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*date = EntityDate(nullTime.Time)
	fmt.Println(value, nullTime.Time)
	return
}

func (date EntityDate) Value() (driver.Value, error) {
	t := time.Time(date)
	if t.IsZero() || t.UnixMicro() == 0 {
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
	if len(b) < 1 {
		return nil
	}
	str := string(b)
	if len(str) < 1 {
		return nil
	}
	t, err := strconv.ParseInt(str, 10, 64)
	if t > 0 && err == nil {
		d := time.UnixMilli(t)
		if d.IsZero() {
			return nil
		}
		*date = EntityDate(d)
		return nil
	}
	if strings.Contains(str, "/") {
		str = strings.ReplaceAll(str, "/", "-")
	}
	d, err := time.Parse("2006-01-02 15:04:05", str)
	if err == nil && !d.IsZero() {
		*date = EntityDate(d)
		return nil
	}

	return nil
}

func (date EntityDate) GetTime() time.Time {
	t := time.Time(date)
	return t
}
