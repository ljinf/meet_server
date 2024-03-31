package service

import (
	"context"
	"fmt"
	"github.com/ljinf/meet_server/internal/model"
	"github.com/ljinf/meet_server/internal/web/cache"
	pb "github.com/ljinf/meet_server/pkg/proto/account"
	"go.uber.org/zap"
)

type AccountService interface {
	CreateAccount(info *model.Register) error
	GetAccountInfo(info *model.Register) (*model.Register, error)
	UpdateAccountInfo(info *model.Register) error //修改phone,email,password
	GetUserInfo(userId int64) (*model.UserInfo, error)
	UpdateUserInfo(user *model.UserInfo) error
}

type accountService struct {
	*Service
	ctx   context.Context
	cache cache.AccountCache
}

func NewAccountService(
	srv *Service,
	cache cache.AccountCache,
) AccountService {
	return &accountService{
		Service: srv,
		ctx:     context.Background(),
		cache:   cache,
	}
}

func (srv *accountService) CreateAccount(info *model.Register) error {

	serviceClient, err := srv.getRpcClient()
	if err != nil {
		return err
	}

	req := &pb.CreateAccountReq{
		Phone:    info.Phone,
		Email:    info.Email,
		Password: info.Password,
	}
	_, err = serviceClient.CreateAccount(srv.ctx, req)

	return err
}

func (srv *accountService) GetAccountInfo(info *model.Register) (*model.Register, error) {

	rpcClient, err := srv.getRpcClient()
	if err != nil {
		return nil, err
	}

	accountInfo, err := rpcClient.GetAccountInfo(srv.ctx, &pb.AccountInfoReq{
		Phone: info.Phone,
		Email: info.Email,
	})
	if err != nil {
		return nil, err
	}

	res := model.Register{
		UserId:   accountInfo.UserId,
		Phone:    accountInfo.Phone,
		Email:    accountInfo.Email,
		Password: accountInfo.Password,
		Salt:     accountInfo.Salt,
	}

	return &res, nil
}

func (srv *accountService) UpdateAccountInfo(info *model.Register) error {

	rpcClient, err := srv.getRpcClient()
	if err != nil {
		return err
	}

	_, err = rpcClient.UpdateAccountInfo(srv.ctx, &pb.UpdateAccountInfoReq{
		UserId:   info.UserId,
		Phone:    info.Phone,
		Email:    info.Email,
		Password: info.Password,
	})

	return err
}

func (srv *accountService) GetUserInfo(userId int64) (*model.UserInfo, error) {

	infoCache, err := srv.cache.GetUserInfoCache(userId)
	if err != nil {
		srv.logger.Error(fmt.Sprintf("GetUserInfoCache %+v", err), zap.Any("userId", userId))
	}

	if infoCache != nil {
		return infoCache, nil
	}

	serviceClient, err := srv.getRpcClient()
	if err != nil {
		return nil, err
	}

	req := pb.UserInfoReq{
		UserId: userId,
	}
	result, err := serviceClient.GetUserInfo(srv.ctx, &req)
	if err != nil {
		return nil, err
	}

	userInfo := &model.UserInfo{
		UserId:   result.UserId,
		NickName: result.NickName,
		Avatar:   result.Avatar,
		Gender:   int(result.Gender),
		Online:   result.Online,
	}
	return userInfo, nil
}

func (srv *accountService) UpdateUserInfo(user *model.UserInfo) error {
	rpcClient, err := srv.getRpcClient()
	if err != nil {
		return err
	}

	result, err := rpcClient.UpdateUserInfo(srv.ctx, &pb.UpdateUserInfoReq{
		UserId:   user.UserId,
		Avatar:   user.Avatar,
		NickName: user.NickName,
	})
	if err != nil {
		return err
	}

	userInfo := &model.UserInfo{
		UserId:   result.UserId,
		NickName: result.NickName,
		Avatar:   result.Avatar,
		Gender:   int(result.Gender),
	}

	if err = srv.cache.SetUserInfoCache(userInfo); err != nil {
		srv.logger.Error(fmt.Sprintf("SetUserInfoCache %+v", err), zap.Any("userInfo", user))
	}

	return nil
}

func (srv *accountService) getRpcClient() (pb.AccountServiceClient, error) {
	//var rpcConn *grpc.ClientConn
	//var err error
	//if srv.rpcClient.Addr != "" {
	//	rpcConn, err = srv.rpcClient.Dial()
	//} else {
	//
	//}

	rpcConn, err := srv.rpcClient.DialWithConsul()
	if err != nil {
		return nil, err
	}
	serviceClient := pb.NewAccountServiceClient(rpcConn)
	return serviceClient, nil
}
