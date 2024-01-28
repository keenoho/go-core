package extend

import (
	"net/http"

	"gitee.com/keenoho/go-core"
	"github.com/gin-gonic/gin"
)

func HttpNotFoundMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, status := core.MakeResponse(nil, -1, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		ctx.AbortWithStatusJSON(status, data)
	}
}
