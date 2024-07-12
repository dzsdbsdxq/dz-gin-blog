package service

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts/vo"
)

type IPostService interface {
	CreatePost(post *vo.PostReq) error
	FindPostByPostId(id uint) (*vo.PostsRes, error)
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
func (ps *PostService) CreatePost(post *vo.PostReq) error {
	//TODO 需要在这里处理一些额外的事情
	return ps.repo.CreatePost(post)
}
func (ps *PostService) FindPostByPostId(id uint) (*vo.PostsRes, error) {
	post, err := ps.repo.FindPostById(id)
	if err != nil {
		return nil, err
	}
	return &vo.PostsRes{
		ID:              post.ID,
		Title:           post.Title,
		Body:            post.Body,
		Extend:          post.Extend,
		Thumb:           post.Thumb,
		CommentAccepted: post.CommentAccepted,
		Recommend:       post.Recommend,
		Status:          post.Status,
		Top:             post.Top,
		SortBy:          post.SortBy,
		Flag:            post.Flag,
		Views:           post.Views,
		Likes:           post.Likes,
		Password:        post.Password,
		Slug:            post.Slug,
		Desc:            post.Desc,
		SeoKey:          post.SeoKey,
		SeoDesc:         post.SeoDesc,
		CreatedAt:       post.CreatedAt,
	}, nil

}
