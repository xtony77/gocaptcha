package route

import (
	"github.com/gin-gonic/gin"
	config "gocaptcha/configs"
	"gocaptcha/internal/gin/handler/captcha"
	"gocaptcha/internal/gin/middleware"
)

func SetupRouter(config *config.Config) *gin.Engine {
	if config.Gin.Mode == "RELEASE" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(middleware.InitLogger)
	api := r.Group("/api")
	{
		api.GET("/captcha", captcha.GetCaptcha)
		api.POST("/captcha", captcha.VerifityCaptcha)
	}

	return r
}
