package service

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts_tags/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts_tags/vo"
)

type IPostTagService interface {
	AdminCreatePostTag(pts []*vo.PostTagReq) error
	AdminDeletePostTag(pts []*vo.PostTagReq) error
}

var _ IPostTagService = (*PostTagService)(nil)

type PostTagService struct {
	repo repo.IPostTagRepository
}

func (p *PostTagService) AdminDeletePostTag(pts []*vo.PostTagReq) error {
	return p.repo.AdminDeletePostTag(pts)
}

func (p *PostTagService) AdminCreatePostTag(pts []*vo.PostTagReq) error {
	return p.repo.AdminCreatePostTag(pts)
}

func NewPostTagService(repo repo.IPostTagRepository) *PostTagService {
	s := &PostTagService{
		repo: repo,
	}
	return s
}
