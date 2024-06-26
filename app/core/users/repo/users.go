package repo

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/users/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/users/repo/dao"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/users/vo"
	"github.com/gofrs/uuid/v5"
)

type IUserRepository interface {
	AdminCreateUser(post *vo.UserReq) error
	AdminGetUserInfoById(id uint) *model.SysUsers
}

var _ IUserRepository = (*UserRepository)(nil)

func NewUserRepository(dao dao.IUserDao) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

type UserRepository struct {
	dao dao.IUserDao
}

func (ur *UserRepository) AdminCreateUser(user *vo.UserReq) error {
	return ur.dao.Create(&model.SysUsers{
		Enable:    user.Enable,
		UserName:  user.UserName,
		Password:  user.Password,
		Uuid:      uuid.Must(uuid.NewV4()).String(),
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
		WebLink:   user.WebLink,
		Desc:      user.Desc,
	})
}

func (ur *UserRepository) AdminGetUserInfoById(id uint) *model.SysUsers {
	user, err := ur.dao.GetUserDetailById(id)
	if err != nil {
		return nil
	}
	return user
}
