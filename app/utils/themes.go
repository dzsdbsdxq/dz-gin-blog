package utils

import (
	"errors"
	"fmt"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"io"
	"os"
	"path/filepath"
)

func ReadConfigFile(name string) ([]byte, error) {
	themeBasePath, _ := os.Getwd()
	fmt.Println(themeBasePath)
	if global.G_DZ_CONFIG.System.ThemePath != "" {
		themeBasePath = global.G_DZ_CONFIG.System.ThemePath
	}
	filePath := filepath.Join(themeBasePath, "themes", name, "settings.json")
	fmt.Println(filePath)
	open, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("主题配置文件读取错误")
	}
	defer func(open *os.File) {
		_ = open.Close()
	}(open)
	byteValue, err := io.ReadAll(open)
	if err != nil {
		return nil, err
	}
	return byteValue, nil
}
