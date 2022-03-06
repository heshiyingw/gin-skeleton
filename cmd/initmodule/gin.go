package initmodule

import (
	"github.com/gin-gonic/gin"
	gin_middleware "github.com/heshiyingw/gin-middleware"
	"go.uber.org/zap"
	internalHttp "skeleton/internal/http"
)

// Gin 初始化gin
func Gin() *gin.Engine {
	engine := gin.New()
	logger, _ := zap.NewProductionConfig().Build()

	engine.Use(gin_middleware.GinLogger(logger), gin_middleware.GinRecovery(logger, true))
	internalHttp.SetRouter(engine)
	return engine
}
