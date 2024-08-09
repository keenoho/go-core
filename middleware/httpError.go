package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keenoho/go-core"
)

func HttpErrorMiddleware(errorHandler ...func(any)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				var status int = http.StatusInternalServerError
				var msg string = http.StatusText(status)
				var code int64 = -1
				var data any = nil

				errData, isErrorData := err.(core.ErrorData)
				if isErrorData {
					if errData.Status > 0 {
						status = errData.Status
					}
					if errData.Code != 0 {
						code = errData.Code
						toStatus := core.CODE_STATUS_MAP[errData.Code]
						if toStatus > -1 {
							status = toStatus
						}
					}
					if errData.Msg != "" {
						msg = errData.Msg
					} else {
						codeMsg := core.CODE_MESSAGE_MAP[errData.Code]
						httpMsg := http.StatusText(errData.Status)
						if codeMsg != "" {
							msg = codeMsg
						} else if httpMsg != "" {
							msg = httpMsg
						}

					}
					if errData.Error != nil {
						data = errData.Error
					}
				}

				responseData, status := core.MakeResponse(data, code, msg, status)
				ctx.AbortWithStatusJSON(status, responseData)

				if len(errorHandler) > 0 {
					for _, handler := range errorHandler {
						go handler(err)
					}
				}

				if !isErrorData {
					core.Log("[Error] %v", err)
				}
			}
		}()
		ctx.Next()
	}
}
