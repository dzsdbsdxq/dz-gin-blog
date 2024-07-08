package utils

import (
	"bytes"
	"errors"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"golang.org/x/image/webp"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"strings"
)

// DecodeImageWidthHeight 解析图片的宽高信息
func DecodeImageWidthHeight(imgBytes []byte, fileType string) (int, int, error) {
	var (
		imgConf image.Config
		err     error
	)
	switch strings.ToLower(fileType) {
	case "jpg", "jpeg":
		imgConf, err = jpeg.DecodeConfig(bytes.NewReader(imgBytes))
	case "webp":
		imgConf, err = webp.DecodeConfig(bytes.NewReader(imgBytes))
	case "png":
		imgConf, err = png.DecodeConfig(bytes.NewReader(imgBytes))
	case "tif", "tiff":
		imgConf, err = tiff.DecodeConfig(bytes.NewReader(imgBytes))
	case "gif":
		imgConf, err = gif.DecodeConfig(bytes.NewReader(imgBytes))
	case "bmp":
		imgConf, err = bmp.DecodeConfig(bytes.NewReader(imgBytes))
	default:
		return 0, 0, errors.New("unknown file type")
	}
	if err != nil {
		return 0, 0, err
	}
	return imgConf.Width, imgConf.Height, nil
}
