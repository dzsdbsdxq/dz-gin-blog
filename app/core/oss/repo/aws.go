package repo

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/oss/vo"
	"mime/multipart"
)

type AwsS3 struct{}

func (*AwsS3) UploadFile(file *multipart.FileHeader) (*vo.OssFile, error) {
	return nil, nil
}
func (*AwsS3) DeleteFile(key string) error {
	return nil
}
