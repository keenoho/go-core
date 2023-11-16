package service

import (
	"encoding/json"
	"strconv"

	"github.com/keenoho/go-core"
)

func CreateServiceApp(middlewares ...MicroServiceMiddleware) *MicroService {
	app := &MicroService{
		RouteMap: make(map[string]MicroServiceControllerFunc),
	}

	if len(middlewares) > 0 {
		for _, m := range middlewares {
			app.Use(m)
		}
	}

	return app
}

/**
 * @params: data bytes, code int, msg string
 **/
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
				byteValue, isBytes := v.([]byte)
				if !isBytes {
					byteValue, _ = json.Marshal(v)
				}
				resData.Data = byteValue
				break
			}
		case 1:
			{
				codeValue, _ := v.(int)
				errValue, isErr := v.(error)

				if isErr {
					val, err := strconv.Atoi(errValue.Error())
					if err == nil {
						codeValue = val
					} else {
						resData.Msg = errValue.Error()
					}
				}

				if codeValue > 0 {
					resData.Code = codeValue
					msg, isExist := core.CodeMsgMap[resData.Code]
					if isExist {
						resData.Msg = msg
					}
				}
				break
			}
		case 2:
			{
				msgValue, _ := v.(string)
				errValue, isErr := v.(error)

				if isErr {
					msgValue = errValue.Error()
				}

				if len(msgValue) > 0 {
					resData.Msg = msgValue
				}
				break
			}
		}
	}
	return resData
}
