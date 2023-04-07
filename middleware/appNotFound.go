package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/keenoho/go-core/util"
	"net/http"
)

func AppNotFoundMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, status := util.MakeResponse(nil, -1, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		ctx.AbortWithStatusJSON(status, data)
	}
}
