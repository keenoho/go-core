package extend

import (
	"net/http"

	"gitee.com/keenoho/go-core"
	"github.com/gin-gonic/gin"
)

func HttpCorsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		corsAllowOrigin := core.ConfigGet("CORS_ALLOW_ORIGIN")
		corsAllowMethods := core.ConfigGet("CORS_ALLOW_METHODS")
		corsAllowHeaders := core.ConfigGet("CORS_ALLOW_HEADERS")
		corsMaxAge := core.ConfigGet("CORS_MAX_AGE")

		ctx.Header("Access-Control-Allow-Origin", corsAllowOrigin)
		ctx.Header("Access-Control-Allow-Methods", corsAllowMethods)
		ctx.Header("Access-Control-Allow-Headers", corsAllowHeaders)
		ctx.Header("Access-control-max-age", corsMaxAge)

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}

		ctx.Next()
	}
}
