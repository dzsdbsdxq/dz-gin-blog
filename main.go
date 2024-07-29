package main

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/initialize"
)

// hello main.go
func main() {
	initialize.InitViper()
	initialize.InitZapLogger()
	initialize.InitCache()
	app, initFunc, err := initializeApp()
	if err != nil {
		panic(err)
	}
	initFunc()
	initialize.Run(app)

}
