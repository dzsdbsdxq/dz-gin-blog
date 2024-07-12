package main

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/initialize"
)

func main() {
	initialize.InitViper()
	initialize.InitZapLogger()
	initialize.InitCache()
	app, err := initializeApp()
	if err != nil {
		panic(err)
	}

	initialize.Run(app)

}
