package repo

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/oss/vo"
	"mime/multipart"
)

type AliyunOSS struct{}

func (*AliyunOSS) UploadFile(file *multipart.FileHeader) (*vo.OssFile, error) {
	return nil, nil
}
func (*AliyunOSS) DeleteFile(key string) error {
	return nil
}
