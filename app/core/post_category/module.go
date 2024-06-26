package post_category

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/post_category/service"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/post_category/vo"
)

type (
	Service        = service.IPostCategoryService
	PostCategoryVO = vo.PostCategoryReq
	Module         struct {
		Svc Service
	}
)
