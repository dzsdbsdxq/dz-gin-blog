package service

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/themes/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/themes/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/utils"
	jsoniter "github.com/json-iterator/go"
)

type IThemesService interface {
	GetThemeSettingConfig(name string) (*vo.ThemeSettingConfig, error)
}

var _ IThemesService = (*ThemesService)(nil)

type ThemesService struct {
	repo repo.IThemesRepository
}

func (t *ThemesService) GetThemeSettingConfig(name string) (*vo.ThemeSettingConfig, error) {
	var config = new(vo.ThemeSettingConfig)
	//读取配置文件
	file, err := utils.ReadConfigFile(name)
	if err != nil {
		return nil, err
	}
	//解析JSON
	err = jsoniter.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func NewThemesService(repo repo.IThemesRepository) *ThemesService {
	return &ThemesService{repo: repo}
}
