package repo

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts/repo/dao"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts/vo"
)

type IPostRepository interface {
	AdminCreatePost(post *vo.PostReq) error
}

var _ IPostRepository = (*PostRepository)(nil)

func NewPostRepository(dao dao.IPostDao) *PostRepository {
	return &PostRepository{
		dao: dao,
	}
}

type PostRepository struct {
	dao dao.IPostDao
}

func (pr *PostRepository) AdminCreatePost(post *vo.PostReq) error {

	return pr.dao.Create(&model.SysPosts{
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
		Views:           0,
		Likes:           0,
		Password:        post.Password,
		Slug:            post.Slug,
		Desc:            post.Desc,
		SeoKey:          post.SeoKey,
		SeoDesc:         post.SeoDesc,
	})
}
