package captcha

import (
	"github.com/gin-gonic/gin"
	captchaPkg "github.com/dchest/captcha"
	"gocaptcha/domain"
	"encoding/base64"
	"bytes"
	"fmt"
	"gocaptcha/internal/gin/handler"
	"go.uber.org/zap"
	"strconv"
)

type GetCaptchaResponse struct {
	Image  string `json:"image"`
	ID     string `json:id`
	Digits string `json:digits`
}

func GetCaptcha(c *gin.Context) {
	id := captchaPkg.NewLen(domain.CaptchaDefaultLen)
	digits := captchaPkg.RandomDigits(domain.CaptchaDefaultLen)
	store := captchaPkg.NewMemoryStore(domain.CaptchaCollectNum, domain.CaptchaExpiration)
	store.Set(id, digits)
	captchaPkg.SetCustomStore(store)

	var content bytes.Buffer
	err := captchaPkg.WriteImage(&content, id, domain.CaptchaStdWidth, domain.CaptchaStdHeight)
	if err != nil {
		zap.S().Warn(err)
		handler.Failed(c, domain.ErrorServer)
	}

	base64Img := base64.StdEncoding.EncodeToString(content.Bytes())

	digitsStr := ""
	for _, val := range digits {
		digitsStr += strconv.Itoa(int(val))
	}

	handler.Success(c, GetCaptchaResponse{
		Image:  fmt.Sprintf("data:image/png;base64,%s", base64Img),
		ID:     id,
		Digits: digitsStr,
	})
}
