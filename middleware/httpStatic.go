package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

func HttpStatic(engine *gin.Engine) gin.IRoutes {
	staticDir := os.Getenv("STATIC_DIR")
	staticPath := os.Getenv("STATIC_PATH")
	if len(staticDir) > 0 && len(staticPath) > 0 {
		return engine.Static(staticPath, staticDir)
	}
	return nil
}
