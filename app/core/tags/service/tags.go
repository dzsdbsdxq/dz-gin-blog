package service

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/tags/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/tags/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/tags/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
)

type ITagService interface {
	AdminCreateTag(req *vo.TagsReq) error
	AdminUpdateTag(id uint64, req *vo.TagsReq) error
	AdminDeleteTag(id uint) error
	SelectTags(req *global.PageInfo) ([]*vo.TagsRes, int64, error)
}

var _ ITagService = (*TagService)(nil)

type TagService struct {
	repo repo.ITagsRepository
}

func NewTagService(repo repo.ITagsRepository) *TagService {
	return &TagService{repo: repo}
}

func (t *TagService) SelectTags(req *global.PageInfo) ([]*vo.TagsRes, int64, error) {

	tags, total, err := t.repo.SelectTags(req.PageSize, req.PageNo)
	if err != nil {
		return nil, 0, err
	}
	tagsRes := make([]*vo.TagsRes, 0)
	for i := 0; i < len(tags); i++ {
		tagsRes = append(tagsRes, &vo.TagsRes{
			Nums:  tags[i].Nums,
			ID:    tags[i].ID,
			Name:  tags[i].Name,
			Slug:  tags[i].Slug,
			Thumb: tags[i].Thumb,
		})
	}
	return tagsRes, total, err
}

func (t *TagService) AdminCreateTag(req *vo.TagsReq) error {
	tagModel := &model.SysTags{
		Name:  req.Name,
		Slug:  req.Slug,
		Thumb: req.Thumb,
		Nums:  0,
	}
	return t.repo.CreateTags(tagModel)
}

func (t *TagService) AdminUpdateTag(id uint64, req *vo.TagsReq) error {
	tagModel := &model.SysTags{
		Model: global.Model{
			ID: id,
		},
		Name:  req.Name,
		Slug:  req.Slug,
		Thumb: req.Thumb,
	}
	return t.repo.UpdateTags(tagModel)
}

func (t *TagService) AdminDeleteTag(id uint) error {
	return t.repo.DeleteTags(uint64(id))
}
