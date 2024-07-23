package repo

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/stat/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/stat/repo/dao"
)

type IStatRepository interface {
	Create(st []*model.SysStat) error
	Update(st *model.SysStat) error
	Select() (sts []*model.SysStat, err error)
	FindByKey(key string) (st *model.SysStat, err error)
}

var _ IStatRepository = (*StatRepository)(nil)

func NewStatRepository(dao dao.IStatDao) *StatRepository {
	return &StatRepository{
		dao: dao,
	}
}

type StatRepository struct {
	dao dao.IStatDao
}

func (s *StatRepository) FindByKey(key string) (st *model.SysStat, err error) {
	return s.dao.FindByKey(key)
}

func (s *StatRepository) Select() (sts []*model.SysStat, err error) {
	return s.dao.Select()
}

func (s *StatRepository) Create(st []*model.SysStat) error {
	return s.dao.Create(st)
}

func (s *StatRepository) Update(st *model.SysStat) error {
	return s.dao.Update(st)
}
