package service

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/stat/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/stat/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/stat/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"strconv"
)

type IStatService interface {
	Create() error
	Update(sts []*vo.StatVO) error
	FindByKey(key string) (*vo.StatRes, error)
	AdminSelect() (sts []*vo.StatVO, err error)
	Select() (sts []*vo.StatRes, err error)
}

var _ IStatService = (*StatService)(nil)

type StatService struct {
	repo repo.IStatRepository
}

func NewStatService(repo repo.IStatRepository) *StatService {
	return &StatService{
		repo: repo,
	}
}

func (s *StatService) Create() error {
	return s.repo.Create(vo.InitializeData())
}

func (s *StatService) Update(sts []*vo.StatVO) error {
	return s.repo.Update(&model.SysStat{
		Model: global.Model{},
		Key:   "",
		Val:   0,
	})
}

func (s *StatService) FindByKey(key string) (*vo.StatRes, error) {
	stm, err := s.repo.FindByKey(key)
	if err != nil {
		return nil, err
	}
	return &vo.StatRes{
		Key: stm.Key,
		Val: strconv.Itoa(int(stm.Val)),
	}, nil
}

func (s *StatService) AdminSelect() (sts []*vo.StatVO, err error) {
	settings, err := s.repo.Select()
	settingVOS := make([]*vo.StatVO, 0)
	for _, st := range settings {
		settingVOS = append(settingVOS, &vo.StatVO{
			Key: st.Key,
			Val: strconv.Itoa(int(st.Val)),
		})
	}
	return settingVOS, err
}

func (s *StatService) Select() (sts []*vo.StatRes, err error) {
	settings, err := s.repo.Select()
	settingRes := make([]*vo.StatRes, 0)
	for _, st := range settings {
		settingRes = append(settingRes, &vo.StatRes{
			Key: st.Key,
			Val: strconv.Itoa(int(st.Val)),
		})
	}
	return settingRes, err
}
