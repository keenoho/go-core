package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func HttpCorsMiddleware(headers ...map[string]string) gin.HandlerFunc {
	corsAllowOrigin := os.Getenv("CORS_ALLOW_ORIGIN")
	corsAllowMethods := os.Getenv("CORS_ALLOW_METHODS")
	corsAllowHeaders := os.Getenv("CORS_ALLOW_HEADERS")
	corsMaxAge := os.Getenv("CORS_MAX_AGE")

	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", corsAllowOrigin)
		ctx.Header("Access-Control-Allow-Methods", corsAllowMethods)
		ctx.Header("Access-Control-Allow-Headers", corsAllowHeaders)
		ctx.Header("Access-control-max-age", corsMaxAge)

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}

		if len(headers) > 0 {
			for _, headerMap := range headers {
				for k, v := range headerMap {
					ctx.Header(k, v)
				}
			}
		}

		ctx.Next()
	}
}
