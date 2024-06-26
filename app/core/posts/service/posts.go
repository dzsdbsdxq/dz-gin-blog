package service

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts/vo"
)

type IPostService interface {
	AdminCreatePost(post *vo.PostReq) error
}

var _ IPostService = (*PostService)(nil)

type PostService struct {
	repo repo.IPostRepository
	//cfgService website_config.Service
	//eventBus   *eventbus.EventBus
}

func NewPostService(repo repo.IPostRepository) *PostService {
	s := &PostService{
		repo: repo,
		//cfgService: cfgService,
		//eventBus:   eventBus,
	}
	//go s.subscribeCommentEvent()
	return s
}
func (ps *PostService) AdminCreatePost(post *vo.PostReq) error {
	//TODO 需要在这里处理一些额外的事情
	return ps.repo.AdminCreatePost(post)
}
