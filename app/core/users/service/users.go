package service

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/users/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/users/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/users/vo"
)

type IUserService interface {
	AdminCreateUser(user *vo.UserReq) error
	AdminUpdateUser(user *vo.UserReq) error
	AdminGetUserInfoById(id uint) *model.SysUsers
}

var _ IUserService = (*UserService)(nil)

type UserService struct {
	repo repo.IUserRepository
	//cfgService website_config.Service
	//eventBus   *eventbus.EventBus
}

func NewUserService(repo repo.IUserRepository) *UserService {
	s := &UserService{
		repo: repo,
		//cfgService: cfgService,
		//eventBus:   eventBus,
	}
	//go s.subscribeCommentEvent()
	return s
}
func (ps *UserService) AdminCreateUser(user *vo.UserReq) error {
	//TODO 需要在这里处理一些额外的事情
	return ps.repo.AdminCreateUser(user)
}

func (ps *UserService) AdminUpdateUser(user *vo.UserReq) error {
	//TODO 需要在这里处理一些额外的事情
	return ps.repo.AdminCreateUser(user)
}

func (ps *UserService) AdminGetUserInfoById(id uint) *model.SysUsers {
	//TODO 需要在这里处理一些额外的事情
	return ps.repo.AdminGetUserInfoById(id)
}
