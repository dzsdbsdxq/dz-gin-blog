package service

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/setting/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/setting/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/setting/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
)

type ISettingService interface {
	Create() error
	Update(sts []*vo.SettingReq) error
	FindByKey(key string) (*vo.SettingRes, error)
	AdminSelect() (sts []*vo.SettingVO, err error)
	Select() (sts []*vo.SettingRes, err error)
}

var _ ISettingService = (*SettingService)(nil)

type SettingService struct {
	repo repo.ISettingRepository
}

func (s *SettingService) Select() (sts []*vo.SettingRes, err error) {
	settings, err := s.repo.Select()
	settingRes := make([]*vo.SettingRes, 0)
	for _, st := range settings {
		settingRes = append(settingRes, &vo.SettingRes{
			Key: st.Key,
			Val: st.Val,
		})
	}
	return settingRes, err
}

func (s *SettingService) Update(sts []*vo.SettingReq) error {
	return s.repo.Update(&model.SysSetting{
		Model:    global.Model{},
		Required: 0,
		Key:      "",
		Val:      "",
		Name:     "",
		Desc:     "",
		Type:     "",
		Comp:     "",
	})
}

func (s *SettingService) FindByKey(key string) (*vo.SettingRes, error) {
	stm, err := s.repo.FindByKey(key)
	if err != nil {
		return nil, err
	}
	return &vo.SettingRes{
		Key: stm.Key,
		Val: stm.Val,
	}, nil
}

func (s *SettingService) AdminSelect() (sts []*vo.SettingVO, err error) {
	settings, err := s.repo.Select()
	settingVOS := make([]*vo.SettingVO, 0)
	for _, st := range settings {
		settingVOS = append(settingVOS, &vo.SettingVO{
			Required: st.Required,
			Key:      st.Key,
			Val:      st.Val,
			Name:     st.Name,
			Desc:     st.Desc,
			Type:     st.Type,
			Comp:     st.Comp,
		})
	}
	return settingVOS, err
}

func (s *SettingService) Create() error {
	return s.repo.Create(vo.InitializeData())
}

func NewSettingService(repo repo.ISettingRepository) *SettingService {
	return &SettingService{
		repo: repo,
	}
}
