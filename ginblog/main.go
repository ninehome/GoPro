package main

import (
	"ginblog/model"
	"ginblog/routes"
	"ginblog/utils"
)

func main() {
	//引用数据库
	model.InitDB()
	engin := routes.InitRouter()
	engin.Run(utils.HttpPort)

}
