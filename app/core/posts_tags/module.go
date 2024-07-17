package posts_tags

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts_tags/model"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts_tags/service"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/posts_tags/vo"
)

type (
	Service      = service.IPostTagService
	PostTagVO    = vo.PostTagReq
	PostTagModel = model.SysPostTag
	Module       struct {
		Svc Service
	}
)
