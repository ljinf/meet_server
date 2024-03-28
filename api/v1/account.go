package v1

type RegisterRequest struct {
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required" example:"123456"`
}

type LoginRequest struct {
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required" example:"123456"`
}

type LoginResponseData struct {
	AccessToken string `json:"accessToken"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type UpdateAccountReq struct {
	UserId   int64  `json:"user_id" binding:"required"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserInfoRequest struct {
	UserId   int64  `json:"user_id" binding:"required"`
	Avatar   string `json:"avatar"` //头像
	Nickname string `json:"nickname" example:"alan"`
}

type GetProfileResponseData struct {
	UserId   int64  `json:"user_id"`
	NickName string `json:"nick_name"` //昵称
	Avatar   string `json:"avatar"`    //头像
	Gender   int    `json:"gender"`    //性别
}
