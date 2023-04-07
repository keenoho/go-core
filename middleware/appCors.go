package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/keenoho/go-core/common"
	"net/http"
)

func AppCorsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		conf := common.GetConfig()

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
