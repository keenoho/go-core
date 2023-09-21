package service

import (
	"encoding/json"

	"github.com/keenoho/go-core"
)

func MakeResponse(args ...any) ResponseData {
	resData := ResponseData{
		Data: nil,
		Code: 0,
		Msg:  "ok",
	}
	for i, v := range args {
		switch i {
		case 0:
			{
				resData.Data = v
				break
			}
		case 1:
			{
				resData.Code = v.(int)
				msg, isExist := core.CodeMsgMap[int(resData.Code)]
				if isExist {
					resData.Msg = msg
				}
				break
			}
		case 2:
			{
				if len(v.(string)) > 0 {
					resData.Msg = v.(string)
				}
				break
			}
		}
	}
	return resData
}

func DataToBytes(data any) ([]byte, error) {
	dataByte, err := json.Marshal(data)
	return dataByte, err
}
