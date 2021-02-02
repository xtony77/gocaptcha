package domain

type ErrorFormat struct {
	Code    int
	Message string
}

var (
	ErrorServer     = ErrorFormat{Code: 500, Message: "Server Error"}
	ErrorBadRequest = ErrorFormat{Code: 400, Message: "bad request"}
	ErrorCaptcha = ErrorFormat{Code: 4001, Message: "error captcha"}
)
