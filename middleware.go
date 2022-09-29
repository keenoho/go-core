package core

import (
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

// 404页面
func NotFoundHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		conf := GetConfig()
		data, status := MakeResponse(nil, -1, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		if conf.ResponseMode == RESPONSE_MODE_ALWAYS200 {
			status = 200
		}
		ctx.JSON(status, data)
	}
}

// 错误捕获
func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				conf := GetConfig()
				status := http.StatusInternalServerError
				msg := http.StatusText(status)
				code := -1
				var data any

				if reflect.TypeOf(err).Name() == "ErrorData" {
					if err.(ErrorData).Status > 0 {
						status = err.(ErrorData).Status
					}
					if err.(ErrorData).Code != 0 {
						code = err.(ErrorData).Code
						toStatus := CodeStatusMap[err.(ErrorData).Code]
						if toStatus > -1 {
							status = toStatus
						}
					}
					if err.(ErrorData).Msg != "" {
						msg = err.(ErrorData).Msg
					} else {
						codeMsg := CodeMsgMap[err.(ErrorData).Code]
						httpMsg := http.StatusText(err.(ErrorData).Status)
						if codeMsg != "" {
							msg = codeMsg
						} else if httpMsg != "" {
							msg = httpMsg
						}

					}
					if err.(ErrorData).Error != nil {
						data = err.(ErrorData).Error
					}
				}

				if conf.ResponseMode == RESPONSE_MODE_ALWAYS200 {
					status = 200
				}

				responseData, status := MakeResponse(data, code, msg, status)
				ctx.AbortWithStatusJSON(status, responseData)
				log.Println("Error:", err)
			}
		}()
		ctx.Next()
	}
}

// 跨域
func CorsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		conf := GetConfig()

		// 处理cors
		ctx.Header("Access-Control-Allow-Origin", conf.CorsAllowOrigin)
		ctx.Header("Access-Control-Allow-Methods", conf.CorsAllowMethods)
		ctx.Header("Access-Control-Allow-Headers", conf.CorsAllowHeaders)
		ctx.Header("Access-control-max-age", "3600")

		// 处理options
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}

		ctx.Next()
	}
}

// 日志
func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}

// 授权session
func SessionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerApp := ctx.GetHeader("x-app")
		headerSignature := ctx.GetHeader("x-signature")

		// 检查判空,解析
		if headerApp != "" && headerSignature != "" {
			data, err := ParseSignature(headerSignature, headerApp)
			if err == nil {
				ctx.Set("signatureData", data)
			}
		}
		ctx.Next()
	}
}
