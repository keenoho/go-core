package service

import "github.com/keenoho/go-core"

func MakeResponse(args ...any) (ResponseData, bool) {
	toJsonString := true
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
				resData.Msg = v.(string)
				break
			}
		}
	}
	return resData, toJsonString
}
