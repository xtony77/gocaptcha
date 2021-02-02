package captcha

import (
	"github.com/gin-gonic/gin"
	"gocaptcha/domain"
	"gocaptcha/internal/gin/handler"
	"go.uber.org/zap"
	captchaPkg "github.com/dchest/captcha"
)

type VerifityCaptchaRequest struct {
	ID      string `json:"id" binding:"required"`
	Captcha string `json:"captcha" binding:"required"`
}

func VerifityCaptcha(c *gin.Context) {
	var req VerifityCaptchaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.S().Infof("err: %v req: %v", err, c.Request)
		handler.Failed(c, domain.ErrorBadRequest)
		return
	}

	if !captchaPkg.VerifyString(req.ID, req.Captcha) {
		handler.Failed(c, domain.ErrorCaptcha)
		return
	}

	handler.Success(c, gin.H{})
}
