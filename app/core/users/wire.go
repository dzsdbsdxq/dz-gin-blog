//go:build wireinject

package users

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/users/controller"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/users/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/users/repo/dao"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/users/service"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var UsersProviders = wire.NewSet(
	controller.NewUserHandler,
	service.NewUserService,
	repo.NewUserRepository,
	dao.NewUserDao,
	wire.Bind(new(service.IUserService), new(*service.UserService)),
	wire.Bind(new(repo.IUserRepository), new(*repo.UserRepository)),
	wire.Bind(new(dao.IUserDao), new(*dao.UserDao)),
)

func InitUsersModule(db *gorm.DB) *Module {
	panic(wire.Build(
		UsersProviders,
		wire.Struct(new(Module), "Svc", "Hdl"),
	))
}
