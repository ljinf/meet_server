package repository

import (
	"github.com/ljinf/meet_server/internal/model"
	"strings"
)

type AccountRepository interface {
	CreateAccount(info *model.Register) error
	GetAccountInfoById(userId int64) (*model.Register, error)
	GetAccountInfo(phone, email string) (*model.Register, error)
	UpdateAccountInfo(info *model.Register) error
	CreateUserInfo(user *model.UserInfo) error
	GetUserInfo(userId int64) (*model.UserInfo, error)
	UpdateUserInfo(info *model.UserInfo) error
}

type accountRepository struct {
	*Repository
}

func NewAccountRepository(repo *Repository) AccountRepository {
	return &accountRepository{
		Repository: repo,
	}
}

func (a *accountRepository) CreateAccount(info *model.Register) error {
	return a.db.Create(info).Error
}

func (a *accountRepository) GetAccountInfo(phone, email string) (*model.Register, error) {

	var conds []string
	if phone != "" {
		conds = append(conds, "phone =? ")
	}
	if email != "" {
		conds = append(conds, "email =? ")
	}

	var accountInfo model.Register
	err := a.db.Where(strings.Join(conds, " and ")).First(&accountInfo).Error

	return &accountInfo, err

}

func (a *accountRepository) GetAccountInfoById(userId int64) (*model.Register, error) {
	var accountInfo model.Register
	err := a.db.Where("user_id=?", userId).First(&accountInfo).Error
	return &accountInfo, err
}

func (a *accountRepository) UpdateAccountInfo(info *model.Register) error {
	return a.db.Save(info).Error
}

func (a *accountRepository) CreateUserInfo(user *model.UserInfo) error {
	return a.db.Save(user).Error
}

func (a *accountRepository) GetUserInfo(userId int64) (*model.UserInfo, error) {
	var userInfo model.UserInfo
	err := a.db.Where("user_id=?", userId).First(&userInfo).Error
	return &userInfo, err
}

func (a *accountRepository) UpdateUserInfo(info *model.UserInfo) error {
	return a.db.Where("user_id=?", info.UserId).Save(info).Error
}
