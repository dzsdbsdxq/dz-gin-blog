package service

import "github.com/dzsdbsdxq/dz-gin-blog/app/core/comments/repo"

type ICommentService interface {
}
type CommentService struct {
	repo repo.ICommentsRepository
}

var _ ICommentService = (*CommentService)(nil)

func NewCommentService(repo repo.ICommentsRepository) *CommentService {
	return &CommentService{
		repo: repo,
	}
}
