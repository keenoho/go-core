package core

import (
	"github.com/gin-gonic/gin"
)

type Context struct {
	App                *App
	HttpServiceContext *gin.Context
}
