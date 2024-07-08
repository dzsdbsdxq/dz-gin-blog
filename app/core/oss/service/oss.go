package service

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/oss/repo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/oss/vo"
	"mime/multipart"
)

type IOssService interface {
	UploadFile(file *multipart.FileHeader) (*vo.OssFile, error)
	DeleteFile(key string) error
}

var _ IOssService = (*OssService)(nil)

type OssService struct {
	repo repo.IOssRepository
}

func NewOssService(repo repo.IOssRepository) *OssService {
	return &OssService{
		repo: repo,
	}
}

func (o *OssService) UploadFile(file *multipart.FileHeader) (*vo.OssFile, error) {
	uploadFile, err := o.repo.UploadFile(file)
	if err != nil {
		return nil, err
	}
	return uploadFile, nil
}

func (o *OssService) DeleteFile(key string) error {
	return o.repo.DeleteFile(key)
}
