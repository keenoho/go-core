package core

import (
	"fmt"
	"strconv"
	"time"
)

func AddCodeMsgMap(msgMap map[int]string) {
	for key, value := range msgMap {
		CodeMsgMap[key] = value
	}
}

func AddCodeStatusMap(statusMap map[int]int) {
	for key, value := range statusMap {
		CodeStatusMap[key] = value
	}
}

func GetStartUpAddress() string {
	conf := GetConfig()
	return fmt.Sprintf("%s:%s", conf["Host"], conf["Port"])
}

func GetRegisterAddress() string {
	conf := GetConfig()
	return fmt.Sprintf("%s:%s", conf["RegisterHost"], conf["RegisterPort"])
}

/**
 * @params: data any, code int, msg string, status int
 **/
func MakeResponse(args ...any) (ResponseData, int) {
	now := time.Now()
	status := 200
	resData := ResponseData{
		Data: nil,
		Code: 0,
		Msg:  "ok",
		Time: now.UnixMilli(),
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
					msg, isExist := CodeMsgMap[resData.Code]
					if isExist {
						resData.Msg = msg
					}
					stu, isExist := CodeStatusMap[resData.Code]
					if isExist {
						status = stu
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
		case 3:
			{
				statusValue, isInt := v.(int)

				if isInt {
					status = statusValue
				}
				break
			}
		}
	}
	return resData, status
}
