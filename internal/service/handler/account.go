package handler

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/ljinf/meet_server/internal/err_msg"
	"github.com/ljinf/meet_server/internal/model"
	"github.com/ljinf/meet_server/internal/service/repository"
	pwd_encoder "github.com/ljinf/meet_server/pkg/helper/pwd-encoder"
	"github.com/ljinf/meet_server/pkg/proto/account"
	"go.uber.org/zap"
)

type AccountServerHandler struct {
	*Handler
	repo repository.AccountRepository
}

func NewAccountServerHandler(server *Handler, repo repository.AccountRepository) *AccountServerHandler {
	return &AccountServerHandler{
		Handler: server,
		repo:    repo,
	}
}

func (as *AccountServerHandler) CreateAccount(ctx context.Context, req *account.CreateAccountReq) (*account.CreateAccountRes, error) {

	userId, err := as.sid.GenUint64()
	if err != nil {
		as.logger.Error(err.Error())
		return nil, err_msg.ErrCreateIdFailed
	}

	salt, encodePwd := pwd_encoder.PwdEncode(req.GetPassword())
	info := model.Register{
		UserId:   int64(userId),
		Phone:    req.Phone,
		Email:    req.Email,
		Password: encodePwd,
		Salt:     salt,
	}

	if err := as.repo.CreateAccount(&info); err != nil {
		as.logger.Error(err.Error(), zap.String("phone", req.Phone),
			zap.String("email", req.Email))
		return nil, err_msg.ErrCreateUserFailed
	}

	if info.UserId == 0 {
		as.logger.Error("err:create account return userId=0", zap.String("phone", req.Phone),
			zap.String("email", req.Email))
		return nil, err_msg.ErrCreateUserFailed
	}

	//创建新用户信息
	userInfo := model.UserInfo{
		UserId: info.UserId,
		Online: false,
		Status: 1,
	}
	if err = as.repo.CreateUserInfo(&userInfo); err != nil {
		as.logger.Error(fmt.Sprintf("初始化用户信息失败 %v", err.Error()), zap.Any("user", userId))
		return nil, err_msg.ErrCreateUserInfoFailed
	}

	return &account.CreateAccountRes{
		UserId: info.UserId,
		Phone:  info.Phone,
		Email:  info.Email,
	}, nil

}

func (as *AccountServerHandler) GetAccountInfo(ctx context.Context, req *account.AccountInfoReq) (*account.AccountInfoRes, error) {
	accountInfo, err := as.repo.GetAccountInfo(req.GetPhone(), req.GetEmail())
	if err != nil {
		as.logger.Error(err.Error(), zap.String("phone", req.Phone), zap.String("email", req.Phone))
		return nil, err_msg.ErrInternalServerError
	}

	return &account.AccountInfoRes{
		UserId:   accountInfo.UserId,
		Phone:    accountInfo.Phone,
		Email:    accountInfo.Email,
		Password: accountInfo.Password,
		Salt:     accountInfo.Salt,
	}, nil
}

func (as *AccountServerHandler) UpdateAccountInfo(ctx context.Context, req *account.UpdateAccountInfoReq) (*empty.Empty, error) {

	infoById, err := as.repo.GetAccountInfoById(req.UserId)
	if err != nil {
		as.logger.Error(err.Error(), zap.Any("userId", req.UserId))
		return nil, err_msg.ErrInternalServerError
	}

	//检查新的phone，email是否已被绑定
	if req.Phone != "" && req.Phone != infoById.Phone {
		infoByPhone, err := as.repo.GetAccountInfo(req.Phone, "")
		if err != nil {
			as.logger.Error(err.Error(), zap.Any("phone", req.Phone))
			return nil, err_msg.ErrInternalServerError
		}

		if infoByPhone != nil {
			return nil, err_msg.ErrPhoneAlreadyUse
		}

	}

	if req.Email != "" && req.Email != infoById.Email {
		infoByEmail, err := as.repo.GetAccountInfo("", req.Email)
		if err != nil {
			as.logger.Error(err.Error(), zap.Any("email", req.Email))
			return nil, err_msg.ErrInternalServerError
		}

		if infoByEmail != nil {
			return nil, err_msg.ErrEmailAlreadyUse
		}
	}

	//正常修改信息
	accountInfo := model.Register{
		UserId: req.UserId,
		Phone:  req.Phone,
		Email:  req.Email,
	}

	if req.Password != "" {
		salt, encodePwd := pwd_encoder.PwdEncode(req.GetPassword())
		accountInfo.Password = encodePwd
		accountInfo.Salt = salt
	}

	if err := as.repo.UpdateAccountInfo(&accountInfo); err != nil {
		as.logger.Error(err.Error(), zap.Any("accountInfo", accountInfo))
		return nil, err_msg.ErrUpdateUserInfoFailed
	}

	return &empty.Empty{}, nil

}

func (as *AccountServerHandler) GetUserInfo(ctx context.Context, req *account.UserInfoReq) (*account.UserInfoRes, error) {
	userInfo, err := as.repo.GetUserInfo(req.UserId)
	if err != nil {
		as.logger.Error(err.Error(), zap.Any("userInfo", userInfo))
		return nil, err_msg.ErrInternalServerError
	}

	res := &account.UserInfoRes{
		UserId:   userInfo.UserId,
		NickName: userInfo.NickName,
		Avatar:   userInfo.Avatar,
		Gender:   int32(userInfo.Gender),
		Online:   userInfo.Online,
	}

	return res, nil

}

func (as *AccountServerHandler) UpdateUserInfo(ctx context.Context, req *account.UpdateUserInfoReq) (*account.UpdateUserInfoRes, error) {
	info := model.UserInfo{
		UserId:   req.UserId,
		Avatar:   req.Avatar,
		NickName: req.NickName,
	}
	if err := as.repo.UpdateUserInfo(&info); err != nil {
		as.logger.Error(err.Error(), zap.Any("userInfo", info))
		return nil, err_msg.ErrUpdateUserInfoFailed
	}

	return &account.UpdateUserInfoRes{
		UserId:   req.UserId,
		Avatar:   req.Avatar,
		NickName: req.NickName,
	}, nil
}
