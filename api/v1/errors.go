package v1

var (
	// common errors
	ErrSuccess             = newError(200, "ok")
	ErrBadRequest          = newError(400, "Bad Request")
	ErrUnauthorized        = newError(401, "Unauthorized")
	ErrNotFound            = newError(404, "Not Found")
	ErrInternalServerError = newError(500, "Internal Server Error")
	ErrServerBusynessError = newError(501, "Internal Server Busyness")

	// more biz errors
	ErrParamError           = newError(1001, "参数错误")
	ErrRegisterFailed       = newError(1002, "注册失败")
	ErrGetAccountInfoFailed = newError(1003, "获取注册信息失败")
	ErrEmailAlreadyUse      = newError(1004, "The email is already in use.")
	ErrPhoneAlreadyUse      = newError(1005, "The phone is already in use.")
	ErrUpdateUserInfoFailed = newError(1006, "更新信息失败")
)
