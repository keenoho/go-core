package middleware

import (
	"github.com/gin-gonic/gin"
	core "github.com/keenoho/go-core"
	"github.com/keenoho/go-core/common"
	"github.com/keenoho/go-core/util"
	"log"
	"net/http"
	"reflect"
	"strings"
)

func AppErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				status := http.StatusInternalServerError
				msg := http.StatusText(status)
				code := -1
				var data any

				errName := reflect.TypeOf(err).Name()
				if strings.Index(errName, "ErrorData") > -1 {
					if err.(core.ErrorData).Status > 0 {
						status = err.(core.ErrorData).Status
					}
					if err.(core.ErrorData).Code != 0 {
						code = err.(core.ErrorData).Code
						toStatus := common.CodeStatusMap[err.(core.ErrorData).Code]
						if toStatus > -1 {
							status = toStatus
						}
					}
					if err.(core.ErrorData).Msg != "" {
						msg = err.(core.ErrorData).Msg
					} else {
						codeMsg := common.CodeMsgMap[err.(core.ErrorData).Code]
						httpMsg := http.StatusText(err.(core.ErrorData).Status)
						if codeMsg != "" {
							msg = codeMsg
						} else if httpMsg != "" {
							msg = httpMsg
						}

					}
					if err.(core.ErrorData).Error != nil {
						data = err.(core.ErrorData).Error
					}
				}

				responseData, status := util.MakeResponse(data, code, msg, status)
				ctx.AbortWithStatusJSON(status, responseData)
				log.Println("Error:", err)
			}
		}()
		ctx.Next()
	}
}
