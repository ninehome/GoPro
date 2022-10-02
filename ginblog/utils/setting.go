package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)
var key_db = "database"
var key_server = "server"

// 读取配置文件
func init() {
	//初始GoINI对象
	file, error := ini.Load("config/config.ini")
	if error != nil {
		fmt.Println("配置文件错误：", error)
	}

	//初始化配置参数
	loadServer(file)
	loadData(file)
}

func loadServer(file *ini.File) {
	var Section = file.Section(key_server)
	AppMode = Section.Key("AppMode").MustString("debug")
	HttpPort = Section.Key("HttpPort").MustString(":3000")
}

func loadData(file *ini.File) {
	var Section = file.Section(key_db)
	Db = Section.Key("Db").MustString("test")
	DbHost = Section.Key("DbHost").MustString("localhost")
	DbPort = Section.Key("DbPort").MustString("3306")
	DbUser = Section.Key("DbUser").MustString("ginblog")
	DbPassWord = Section.Key("DbPassWord").MustString("123qwe")
	DbName = Section.Key("DbName").MustString("ginblog")
}
