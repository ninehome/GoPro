package main

import (
	"ginblog/routes"
	"ginblog/utils"
)

func main() {
	engin := routes.InitRouter()
	engin.Run(utils.HttpPort)
}
