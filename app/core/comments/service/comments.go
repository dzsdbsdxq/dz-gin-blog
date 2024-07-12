package service

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/comments/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/comments/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/comments/vo"
)

type ICommentService interface {
	AddComment(req vo.CommentsReq) error
}
type CommentService struct {
	repo repo.ICommentsRepository
}

func (c *CommentService) AddComment(req vo.CommentsReq) error {
	commentModel := &model.SysComments{
		UserType:  req.UserType,
		CommentId: req.CommentId,
		Status:    req.Status,
		Type:      req.Type,
		FromId:    req.FromId,
		NickName:  req.NickName,
		Email:     req.Email,
		Link:      req.Link,
		Body:      req.Body,
		FromTitle: req.FromTitle,
	}
	return c.repo.CreateComments(commentModel)
}

var _ ICommentService = (*CommentService)(nil)

func NewCommentService(repo repo.ICommentsRepository) *CommentService {
	return &CommentService{
		repo: repo,
	}
}
