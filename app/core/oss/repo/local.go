package repo

import (
	"errors"
	"fmt"
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/oss/vo"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"github.com/dzsdbsdxq/dz-gin-blog/app/utils"
	"go.uber.org/zap"
	"io"
	"mime"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type Local struct{}

func (*Local) UploadFile(file *multipart.FileHeader) (*vo.OssFile, error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	//读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V([]byte(name))
	//拼接新文件名
	fileName := fmt.Sprintf("%s_%s%s", name, time.Now().Format("20060102150405"), ext)
	//尝试创建路径
	mkdirErr := os.MkdirAll(global.G_DZ_CONFIG.Local.StorePath, os.ModePerm)
	if mkdirErr != nil {
		global.G_DZ_LOG.Error("function os.MkdirAll() failed", zap.Any("err", mkdirErr.Error()))
		return nil, errors.New("function os.MkdirAll() failed, err:" + mkdirErr.Error())
	}
	//拼接路径和文件名
	filePath := filepath.Join(global.G_DZ_CONFIG.Local.StorePath, fileName)

	open, openError := file.Open()
	if openError != nil {
		global.G_DZ_LOG.Error("function file.Open() failed", zap.Any("err", openError.Error()))
		return nil, errors.New("function file.Open() failed, err:" + openError.Error())
	}
	defer open.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(filePath)
	if createErr != nil {
		global.G_DZ_LOG.Error("function os.Create() failed", zap.Any("err", createErr.Error()))
		return nil, errors.New("function os.Create() failed, err:" + createErr.Error())
	}
	defer out.Close()
	_, copyErr := io.Copy(out, open) // 传输（拷贝）文件
	if copyErr != nil {
		global.G_DZ_LOG.Error("function io.Copy() failed", zap.Any("err", copyErr.Error()))
		return nil, errors.New("function io.Copy() failed, err:" + copyErr.Error())
	}
	//获取文件大小
	fileInfo, err := out.Stat()
	if err != nil {
		return nil, err
	}
	//w,h,err := utils.DecodeImageWidthHeight(open.)

	return &vo.OssFile{
		FileSize:   uint(fileInfo.Size()),
		FileWidth:  0,
		FileHeight: 0,
		FileName:   file.Filename,
		FileMd5:    name,
		FilePath:   filePath,
		FileExt:    ext,
		FileType:   mime.TypeByExtension(ext),
	}, nil
}
func (*Local) DeleteFile(filePath string) error {
	if err := os.Remove(filePath); err != nil {
		return errors.New("本地文件删除失败, err:" + err.Error())
	}
	return nil
}
