package main

import (
	"github.com/CHN-ChenYi/eLibrary-DB2021/conf"
	"github.com/CHN-ChenYi/eLibrary-DB2021/controller"
	"github.com/CHN-ChenYi/eLibrary-DB2021/model"
)

func main() {
	conf.Init()

	model.Connect()
	model.AutoMigrate()

	controller.Init()
}
