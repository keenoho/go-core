package core

import (
	"fmt"
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
				resData.Code = v.(int)
				break
			}
		case 2:
			{
				resData.Msg = v.(string)
				break
			}
		case 3:
			{
				status = v.(int)
				break
			}
		}
	}
	return resData, status
}

func MakeServiceResponse(args ...any) ServiceResponseData {
	resData := ServiceResponseData{
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
				resData.Code = int64(v.(int))
				break
			}
		case 2:
			{
				resData.Msg = v.(string)
				break
			}
		}
	}
	return resData
}
