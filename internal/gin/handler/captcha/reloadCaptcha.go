package captcha

import (
	"github.com/gin-gonic/gin"
	"gocaptcha/domain"
	"gocaptcha/internal/gin/handler"
	"go.uber.org/zap"
	captchaPkg "github.com/dchest/captcha"
	"encoding/base64"
	"bytes"
	"fmt"
)

type ReloadCaptchaRequest struct {
	ID string `json:"id" binding:"required"`
}

type ReloadCaptchaResponse struct {
	Image  string `json:"image"`
	ID     string `json:"id"`
}

func ReloadCaptcha(c *gin.Context) {
	var req ReloadCaptchaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.S().Infof("err: %v req: %v", err, c.Request)
		handler.Failed(c, domain.ErrorBadRequest)
		return
	}

	if !captchaPkg.Reload(req.ID) {
		handler.Failed(c, domain.ErrorBadRequest)
		return
	}

	var content bytes.Buffer
	err := captchaPkg.WriteImage(&content, req.ID, domain.CaptchaStdWidth, domain.CaptchaStdHeight)
	if err != nil {
		zap.S().Warn(err)
		handler.Failed(c, domain.ErrorServer)
	}
	base64Img := base64.StdEncoding.EncodeToString(content.Bytes())

	handler.Success(c, ReloadCaptchaResponse{
		Image:  fmt.Sprintf("data:image/png;base64,%s", base64Img),
		ID:     req.ID,
	})
}
