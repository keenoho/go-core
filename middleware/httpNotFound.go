package extend

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keenoho/go-core"
)

func HttpNotFoundMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, status := core.MakeResponse(nil, -1, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		ctx.AbortWithStatusJSON(status, data)
	}
}
