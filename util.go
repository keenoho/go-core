package core

import (
	"reflect"
	"strconv"
	"time"
)

func StructConvert(source any, target any, judgeHandler ...func(fieldName string, fieldValue reflect.Value) bool) {
	sourceType := reflect.TypeOf(source)
	sourceValue := reflect.ValueOf(source)
	targetType := reflect.TypeOf(target).Elem()
	targetValue := reflect.ValueOf(target).Elem()

	for i := 0; i < sourceType.NumField(); i++ {
		fieldName := sourceType.Field(i).Name
		fieldValue := sourceValue.FieldByName(fieldName)

		if len(judgeHandler) > 0 {
			if !(judgeHandler[0](fieldName, fieldValue)) {
				continue
			}
		}

		if targetFieldType, ok := targetType.FieldByName(fieldName); ok {
			targetFieldValue := targetValue.FieldByName(fieldName)

			if targetFieldValue.CanSet() && fieldValue.Type() == targetFieldType.Type {
				targetFieldValue.Set(fieldValue)
			}
		}
	}
}

/**
 * @params: data any, code int64, msg string, status int
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
				codeValue, _ := v.(int64)
				codeValueInt, _ := v.(int)

				if codeValue < 1 && codeValueInt > 0 {
					codeValue = int64(codeValueInt)
				}

				errValue, isErr := v.(error)

				if isErr {
					val, err := strconv.ParseInt(errValue.Error(), 10, 64)
					if err == nil {
						codeValue = val
					} else {
						resData.Msg = errValue.Error()
						codeValue = -1
					}
				}

				resData.Code = codeValue
				if codeValue > 0 {
					msg, isExist := CODE_MESSAGE_MAP[resData.Code]
					if isExist {
						resData.Msg = msg
					}
					stu, isExist := CODE_STATUS_MAP[resData.Code]
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
