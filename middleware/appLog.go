package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func AppLoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		latency := time.Since(start)
		log.Printf("[%s][%s] %s - %d %fs %dB", ctx.ClientIP(), ctx.Request.Method, ctx.Request.URL.Path, ctx.Writer.Status(), latency.Seconds(), ctx.Writer.Size())
	}
}
