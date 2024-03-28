package err_msg

import "errors"

var (
	// common errors
	ErrServiceNotFound     = newError(404, "Not Found")
	ErrInternalServerError = newError(500, "Internal Server Error")

	// more biz errors
	ErrCreateIdFailed       = newError(1001, "申请用户ID失败")
	ErrCreateUserFailed     = newError(1002, "创建用户失败")
	ErrCreateUserInfoFailed = newError(1003, "初始化用户消息失败")

	ErrEmailAlreadyUse = newError(1004, "The email is already in use.")
	ErrPhoneAlreadyUse = newError(1005, "The phone is already in use.")

	ErrUpdateUserInfoFailed = newError(1006, "更新信息失败")
)

type Error struct {
	Code    int
	Message string
}

var errorCodeMap = map[error]int{}

func newError(code int, msg string) error {
	err := errors.New(msg)
	errorCodeMap[err] = code
	return err
}
func (e Error) Error() string {
	return e.Message
}
