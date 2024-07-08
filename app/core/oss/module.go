package oss

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/oss/service"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/oss/vo"
)

type (
	OssVO   = vo.OssFile
	Service = service.IOssService
	Module  struct {
		Svc Service
	}
)
