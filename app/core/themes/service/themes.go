package service

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/themes/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/themes/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/themes/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/utils"
	jsoniter "github.com/json-iterator/go"
)

type IThemesService interface {
	Create(themes []*vo.ThemeReq) error
	GetThemeSettingConfig(name string) (*vo.ThemeSettingConfig, error)
	GetThemeSetting(name string) ([]*vo.ThemeKeyValue, error)
}

var _ IThemesService = (*ThemesService)(nil)

type ThemesService struct {
	repo repo.IThemesRepository
}

func (t *ThemesService) Create(themes []*vo.ThemeReq) error {
	thms := make([]*model.SysThemes, 0)
	for _, theme := range themes {
		thms = append(thms, &model.SysThemes{
			Key:  theme.Key,
			Val:  theme.Value,
			Type: theme.Type,
			Name: theme.Desc,
			Desc: theme.Desc,
			Hook: theme.Hook,
		})
	}
	return t.repo.Create(thms)
}

func (t *ThemesService) GetThemeSetting(name string) ([]*vo.ThemeKeyValue, error) {
	themes, err := t.repo.SelectByType(name)
	if err != nil {
		return nil, err
	}
	themesOptions := make([]*vo.ThemeKeyValue, len(themes))
	for _, theme := range themes {
		themesOptions = append(themesOptions, &vo.ThemeKeyValue{
			Value: theme.Val,
			Key:   theme.Key,
		})
	}
	return themesOptions, nil
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
