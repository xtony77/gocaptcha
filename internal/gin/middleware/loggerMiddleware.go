package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gocaptcha/internal/logger"
)

func InitLogger(c *gin.Context) {
	logger := logger.NewLogger()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()
	c.Next()
}
