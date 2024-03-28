package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "github.com/ljinf/meet_server/api/v1"
	"github.com/ljinf/meet_server/internal/model"
	"github.com/ljinf/meet_server/internal/web/service"
	pwd_encoder "github.com/ljinf/meet_server/pkg/helper/pwd-encoder"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type AccountHandler interface {
	CreateAccount(ctx *gin.Context)
	Login(ctx *gin.Context)
	UpdateAccountInfo(ctx *gin.Context) //修改phone,email,password
	GetUserInfo(ctx *gin.Context)
	UpdateUserInfo(ctx *gin.Context)
}

type accountHandler struct {
	*Handler
	accountSrv service.AccountService
}

func NewAccountHandler(h *Handler, service service.AccountService) AccountHandler {
	return &accountHandler{
		Handler:    h,
		accountSrv: service,
	}
}

func (h *accountHandler) CreateAccount(ctx *gin.Context) {
	var param v1.RegisterRequest

	if err := ctx.ShouldBind(&param); err != nil {
		h.logger.Error(err.Error())
		v1.HandleError(ctx, http.StatusOK, v1.ErrBadRequest, nil)
		return
	}

	if (param.Phone == "" && param.Email == "") || param.Password == "" {
		v1.HandleError(ctx, http.StatusOK, v1.ErrParamError, nil)
		return
	}

	info := model.Register{
		Phone:    param.Phone,
		Email:    param.Email,
		Password: param.Password,
	}
	if err := h.accountSrv.CreateAccount(&info); err != nil {
		h.logger.Error(err.Error(), zap.Any("accountSrv.CreateAccount register info", info))
		v1.HandleError(ctx, http.StatusOK, v1.ErrServerBusynessError, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

func (h *accountHandler) Login(ctx *gin.Context) {
	var param v1.RegisterRequest

	if err := ctx.ShouldBind(&param); err != nil {
		h.logger.Error(err.Error())
		v1.HandleError(ctx, http.StatusOK, v1.ErrBadRequest, nil)
		return
	}

	if (param.Phone == "" && param.Email == "") || param.Password == "" {
		v1.HandleError(ctx, http.StatusOK, v1.ErrParamError, nil)
		return
	}

	info := model.Register{
		Phone:    param.Phone,
		Email:    param.Email,
		Password: param.Password,
	}

	accountInfo, err := h.accountSrv.GetAccountInfo(&info)
	if err != nil {
		h.logger.Error(err.Error(), zap.Any("accountSrv.GetAccountInfo", info))
		v1.HandleError(ctx, http.StatusOK, v1.ErrGetAccountInfoFailed, nil)
		return
	}

	if ok := pwd_encoder.PwdDecode(param.Password, accountInfo.Password, accountInfo.Salt); !ok {
		h.logger.Error("登录密码错误", zap.Any("param", fmt.Sprintf("phone:%v email:%v", param.Phone, param.Email)))
		v1.HandleError(ctx, http.StatusOK, v1.ErrParamError, nil)
		return
	}

	token, err := h.jwt.GenToken(accountInfo.UserId, time.Unix(7*24*3600, 0))
	if err != nil {
		h.logger.Error(err.Error(), zap.Any("userId", accountInfo.UserId))
	}

	res := v1.LoginResponseData{
		AccessToken: token,
	}

	v1.HandleSuccess(ctx, res)

}

func (h *accountHandler) UpdateAccountInfo(ctx *gin.Context) {
	var param v1.UpdateAccountReq
	if err := ctx.ShouldBind(&param); err != nil {
		h.logger.Error(fmt.Sprintf("UpdateAccountInfo %+v", err.Error()))
		v1.HandleError(ctx, http.StatusOK, v1.ErrParamError, nil)
		return
	}
	info := model.Register{
		UserId:   param.UserId,
		Phone:    param.Phone,
		Email:    param.Email,
		Password: param.Password,
	}
	if err := h.accountSrv.UpdateAccountInfo(&info); err != nil {
		h.logger.Error(fmt.Sprintf("UpdateAccountInfo %v", err))
		v1.HandleError(ctx, http.StatusOK, v1.ErrServerBusynessError, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

func (h *accountHandler) GetUserInfo(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		h.logger.Error("GetUserInfo 缺失userId", zap.Any("ip", ctx.RemoteIP()))
		v1.HandleError(ctx, http.StatusOK, v1.ErrBadRequest, nil)
		return
	}

	info, err := h.accountSrv.GetUserInfo(userId.(int64))
	if err != nil {
		h.logger.Error(err.Error(), zap.Any("userId", userId))
		v1.HandleError(ctx, http.StatusOK, v1.ErrServerBusynessError, nil)
		return
	}

	v1.HandleSuccess(ctx, info)

}

func (h *accountHandler) UpdateUserInfo(ctx *gin.Context) {
	var param v1.UpdateUserInfoRequest
	if err := ctx.ShouldBind(&param); err != nil {
		h.logger.Error(fmt.Sprintf("UpdateUserInfo %v", err.Error()))
		v1.HandleError(ctx, http.StatusOK, v1.ErrParamError, nil)
		return
	}

	info := model.UserInfo{
		UserId:   param.UserId,
		Avatar:   param.Avatar,
		NickName: param.Nickname,
	}
	if err := h.accountSrv.UpdateUserInfo(&info); err != nil {
		h.logger.Error(err.Error(), zap.Any("user", info))
		v1.HandleError(ctx, http.StatusOK, v1.ErrUpdateUserInfoFailed, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}
