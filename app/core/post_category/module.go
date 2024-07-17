package post_category

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/post_category/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/post_category/service"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/post_category/vo"
)

type (
	Service           = service.IPostCategoryService
	PostCategoryVO    = vo.PostCategoryReq
	PostCategoryModel = model.SysPostCategory
	Module            struct {
		Svc Service
	}
)
