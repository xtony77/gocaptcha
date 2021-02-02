package domain

import "time"

const (
	// Default number of digits in captcha solution.
	CaptchaDefaultLen = 6

	// The number of captchas created that triggers garbage collection used by default store.
	CaptchaCollectNum = 100

	// Expiration time of captchas used by default store.
	CaptchaExpiration = 10 * time.Minute

	// Standard width and height of a captcha image.
	CaptchaStdWidth  = 240
	CaptchaStdHeight = 80
)
