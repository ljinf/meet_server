package model

import (
	"gorm.io/gorm"
)

// 注册表
type Register struct {
	gorm.Model
	UserId   int64  `json:"user_id"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

func (r *Register) TableName() string {
	return "im_register"
}

// 用户信息表
type UserInfo struct {
	gorm.Model

	UserId   int64  `json:"user_id"`
	NickName string `json:"nick_name"` //昵称
	Avatar   string `json:"avatar"`    //头像
	Gender   int    `json:"gender"`    //性别
	Online   bool   `json:"online"`    //是否在线 0:false  1:true
	Status   int    `json:"status"`    //用户状态  0:异常  1:正常
}

func (u *UserInfo) TableName() string {
	return "im_user_info"
}
