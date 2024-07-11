package service

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/post_category/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/post_category/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/post_category/vo"
)

type IPostCategoryService interface {
	CreatePostCategory(pcs []*vo.PostCategoryReq) error
	DeletePostCategory(pcs []*vo.PostCategoryReq) error
	SelectPostCategory(category uint) (pcs []*vo.PostCategoryReq, err error)
	FindPostCategory(category uint, postId uint) (pc *vo.PostCategoryReq, err error)
}

var _ IPostCategoryService = (*PostCategoryService)(nil)

type PostCategoryService struct {
	repo repo.IPostCategoryRepository
}

func (p *PostCategoryService) DeletePostCategory(pcs []*vo.PostCategoryReq) error {
	sysPostCategorys := make([]*model.SysPostCategory, len(pcs))
	for _, pc := range pcs {
		sysPostCategorys = append(sysPostCategorys, &model.SysPostCategory{
			PostId:     pc.PostId,
			CategoryId: pc.Category,
		})
	}
	return p.repo.DeletePostCategory(sysPostCategorys)
}

func (p *PostCategoryService) CreatePostCategory(pcs []*vo.PostCategoryReq) error {
	sysPostCategorys := make([]*model.SysPostCategory, len(pcs))
	for _, pc := range pcs {
		sysPostCategorys = append(sysPostCategorys, &model.SysPostCategory{
			PostId:     pc.PostId,
			CategoryId: pc.Category,
		})
	}
	return p.repo.CreatePostCategory(sysPostCategorys)
}

func (p *PostCategoryService) SelectPostCategory(category uint) (pcs []*vo.PostCategoryReq, err error) {
	postCategorys, err := p.repo.SelectPostCategory(category)
	if err != nil {
		return nil, err
	}
	pcs = make([]*vo.PostCategoryReq, 0)
	for i := 0; i < len(postCategorys); i++ {
		pcs = append(pcs, &vo.PostCategoryReq{
			PostId:   postCategorys[i].PostId,
			Category: postCategorys[i].CategoryId,
		})
	}
	return
}
func (p *PostCategoryService) FindPostCategory(category uint, postId uint) (pc *vo.PostCategoryReq, err error) {
	postCategory, err := p.repo.FindPostCategory(category, postId)
	if err != nil {
		return nil, err
	}
	pc = &vo.PostCategoryReq{
		PostId:   postCategory.PostId,
		Category: postCategory.CategoryId,
	}
	return
}

func NewPostCategoryService(repo repo.IPostCategoryRepository) *PostCategoryService {
	s := &PostCategoryService{
		repo: repo,
	}
	return s
}
