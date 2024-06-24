package initialize

import (
	"fmt"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func InitViper() {
	global.G_DZ_VP = CreateViper("")
}
func CreateViper(path string) *viper.Viper {
	var config string
	var err error
	if path == "" {
		workDir, err := os.Getwd()
		if err != nil {
			panic(fmt.Errorf("读取应用目录失败:%s \n", err))
		}
		config = filepath.Join(workDir, "config.yaml")
	} else {
		config = path
	}
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err = v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.G_DZ_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err = v.Unmarshal(&global.G_DZ_CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
