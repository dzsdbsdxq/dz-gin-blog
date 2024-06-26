package dao

import (
	"errors"
	"fmt"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/users/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"gorm.io/gorm"
)

type IUserDao interface {
	Create(user *model.SysUsers) error
	GetUserDetailById(id uint) (sysUser *model.SysUsers, err error)
	GetUserDetailByUserName(username string) (sysUser *model.SysUsers, err error)
	Update(user *model.SysUsers) error
	DeletePostByIds(ids global.IdsReq) error
}

var _ IUserDao = (*UserDao)(nil)

func NewUserDao() *UserDao {
	return &UserDao{
		coll: global.G_DZ_DB,
	}
}

type UserDao struct {
	coll *gorm.DB
}

func (u *UserDao) Create(user *model.SysUsers) error {
	_, err := u.GetUserDetailByUserName(user.UserName)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("存在相同的用户:%s", user.UserName))
	}
	return u.coll.Create(&user).Error
}

func (u *UserDao) GetUserDetailById(id uint) (sysUser *model.SysUsers, err error) {
	err = u.coll.Where("id = ?", id).First(&sysUser).Error
	return
}

func (u *UserDao) Update(user *model.SysUsers) error {
	_, err := u.GetUserDetailByUserName(user.UserName)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("存在相同的别名:%s", user.UserName))
	}
	return u.coll.Save(&user).Error
}

func (u *UserDao) DeletePostByIds(ids global.IdsReq) error {
	var users []model.SysUsers
	return u.coll.Find(&users, "id in ?", ids.Ids).Delete(&users).Error
}

func (u *UserDao) GetUserDetailByUserName(username string) (sysUser *model.SysUsers, err error) {
	err = u.coll.Where("user_name = ?", username).First(&sysUser).Error
	return
}
