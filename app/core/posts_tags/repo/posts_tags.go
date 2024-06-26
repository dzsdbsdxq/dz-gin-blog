package repo

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts_tags/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts_tags/repo/dao"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts_tags/vo"
)

type IPostTagRepository interface {
	AdminCreatePostTag(pts []*vo.PostTagReq) error
	AdminDeletePostTag(pts []*vo.PostTagReq) error
}

var _ IPostTagRepository = (*PostTagRepository)(nil)

func NewPostTagRepository(dao dao.IPostTagDao) *PostTagRepository {
	return &PostTagRepository{
		dao: dao,
	}
}

type PostTagRepository struct {
	dao dao.IPostTagDao
}

func (p *PostTagRepository) AdminDeletePostTag(pts []*vo.PostTagReq) error {
	sysPostTags := make([]*model.SysPostTag, len(pts))
	for _, pt := range pts {
		sysPostTags = append(sysPostTags, &model.SysPostTag{
			PostId: pt.PostId,
			TagId:  pt.TagId,
		})
	}
	return p.dao.Delete(sysPostTags)
}

func (p *PostTagRepository) AdminCreatePostTag(pts []*vo.PostTagReq) error {
	sysPostTags := make([]*model.SysPostTag, len(pts))
	for _, pt := range pts {
		sysPostTags = append(sysPostTags, &model.SysPostTag{
			PostId: pt.PostId,
			TagId:  pt.TagId,
		})
	}
	return p.dao.Create(sysPostTags)
}
