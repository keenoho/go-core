package core

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"
)

func AppCorsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		conf := GetConfig()

		ctx.Header("Access-Control-Allow-Origin", conf["CorsAllowOrigin"])
		ctx.Header("Access-Control-Allow-Methods", conf["CorsAllowMethods"])
		ctx.Header("Access-Control-Allow-Headers", conf["CorsAllowHeaders"])
		ctx.Header("Access-control-max-age", "3600")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}

		ctx.Next()
	}
}

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

				responseData, status := MakeResponse(data, code, msg, status)
				ctx.AbortWithStatusJSON(status, responseData)
				log.Println("Error:", err)
			}
		}()
		ctx.Next()
	}
}

func AppLoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		latency := time.Since(start)
		log.Printf("[%s][%s] %s - %d %fs %dB", ctx.ClientIP(), ctx.Request.Method, ctx.Request.URL.Path, ctx.Writer.Status(), latency.Seconds(), ctx.Writer.Size())
	}
}

func AppNotFoundMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, status := MakeResponse(nil, -1, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		ctx.AbortWithStatusJSON(status, data)
	}
}
