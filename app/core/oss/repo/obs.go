package repo

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/oss/vo"
	"mime/multipart"
)

type HuaweiObs struct{}

func (*HuaweiObs) UploadFile(file *multipart.FileHeader) (*vo.OssFile, error) {
	return nil, nil
}
func (*HuaweiObs) DeleteFile(key string) error {
	return nil
}
