package service

import "github.com/dzsdbsdxq/dz-gin-blog/app/core/themes/repo"

type IThemesService interface{}

var _ IThemesService = (*ThemesService)(nil)

type ThemesService struct {
	repo repo.IThemesRepository
}

func NewThemesService(repo repo.IThemesRepository) *ThemesService {
	return &ThemesService{repo: repo}
}
