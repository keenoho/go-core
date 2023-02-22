package core

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type ControllerFunc func(ctx *gin.Context) (ResponseData, int)

type ErrorData struct {
	Error  any    `json:"error"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Status int    `json:"status"`
}

type ResponseData struct {
	Data    any    `json:"data"`
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	SysTime int64  `json:"systime"`
}

type SignatureData struct {
	Expired int    `json:"expired"`
	App     int64  `json:"app"`
	Sig     string `json:"sig"`
	Ttl     int    `json:"ttl"`
	Ts      int64  `json:"ts"`
	Nonce   string `json:"nonce"`
	Data    string `json:"data"`
}

type SignatureInnerData struct {
	Id      int64   `json:"id"`
	Account string  `json:"account"`
	Name    string  `json:"name"`
	Role    []int64 `json:"role"`
}

// entity时间
type EntityDate time.Time

func (date *EntityDate) Scan(value interface{}) (err error) {
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
	return time.Date(y, m, d, 0, 0, 0, 0, time.Time(date).Location()), nil
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
