package repo

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/oss/vo"
	"mime/multipart"
)

type TencentCOS struct{}

func (*TencentCOS) UploadFile(file *multipart.FileHeader) (*vo.OssFile, error) {
	return nil, nil
}
func (*TencentCOS) DeleteFile(key string) error {
	return nil
}
