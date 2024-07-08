package repo

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/oss/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"mime/multipart"
)

type IOssRepository interface {
	UploadFile(file *multipart.FileHeader) (*vo.OssFile, error)
	DeleteFile(key string) error
}

var _ IOssRepository = (*OssRepository)(nil)

func NewOssRepository() *OssRepository {
	switch global.G_DZ_CONFIG.System.OssType {
	case "local":
		return &OssRepository{
			dao: &Local{},
		}
	case "qiniu":
		return &OssRepository{
			dao: &Qiniu{},
		}
	case "tencent":
		return &OssRepository{
			dao: &TencentCOS{},
		}
	case "aliyun":
		return &OssRepository{
			dao: &AliyunOSS{},
		}
	case "huawei-obs":
		return &OssRepository{
			dao: &HuaweiObs{},
		}
	case "aws-s3":
		return &OssRepository{
			dao: &AwsS3{},
		}
	default:
		return &OssRepository{
			dao: &Local{},
		}
	}
}

type OssRepository struct {
	dao IOssRepository
}

func (o *OssRepository) UploadFile(file *multipart.FileHeader) (*vo.OssFile, error) {
	return o.dao.UploadFile(file)
}

func (o *OssRepository) DeleteFile(key string) error {
	//TODO implement me
	panic("implement me")
}
