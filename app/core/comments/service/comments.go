package service

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/comments/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/comments/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/comments/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
)

type ICommentService interface {
	AddComment(req vo.CommentsReq) error
	UpdateComment(id uint64, req vo.CommentsReq) error
	DeleteComment(ids global.IdsReq) error
	SelectComments(req vo.GetCommentsListReq) ([]*vo.CommentsRes, int64, error)
	FindCommentById(id uint64) (*vo.CommentsRes, error)
}
type CommentService struct {
	repo repo.ICommentsRepository
}

func (c *CommentService) FindCommentById(id uint64) (*vo.CommentsRes, error) {
	return c.repo.FindCommentById(id)
}

func (c *CommentService) SelectComments(req vo.GetCommentsListReq) ([]*vo.CommentsRes, int64, error) {
	comments, total, err := c.repo.SelectComments(req)
	if err != nil {
		return nil, 0, err
	}
	lists := make([]*vo.CommentsRes, 0)
	for _, comment := range comments {
		lists = append(lists, &vo.CommentsRes{
			UserType:  comment.UserType,
			CommentId: comment.CommentId,
			Status:    comment.Status,
			Type:      comment.Type,
			FromId:    comment.FromId,
			NickName:  comment.NickName,
			Email:     comment.Email,
			Link:      comment.Link,
			Body:      comment.Body,
			FromTitle: comment.FromTitle,
		})
	}
	return lists, total, err
}

func (c *CommentService) DeleteComment(ids global.IdsReq) error {
	return c.repo.DeleteComments(ids)
}

func (c *CommentService) UpdateComment(id uint64, req vo.CommentsReq) error {
	return c.repo.UpdateComments(&model.SysComments{
		Model: global.Model{
			ID: id,
		},
		Status: req.Status,
	})
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
