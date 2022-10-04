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

const load_HOME = "HOME"
const load_WORK = "WORK"

// 读取配置文件
func init() {
	//file, error := ini.Load("config/config.ini")
	file, error := ini.Load("config/configWork.ini")
	if error != nil {
		fmt.Println("配置文件错误：", error)
	}

	//初始化配置参数
	loadData(file, load_WORK)
}

func loadData(file *ini.File, plat string) {

	var Section = file.Section(key_db)
	//公共参数
	AppMode = Section.Key("AppMode").MustString("debug")
	HttpPort = Section.Key("HttpPort").MustString(":3000")

	//数据库参数
	if plat == "HOME" {
		Db = Section.Key("Db").MustString("test")
		DbHost = Section.Key("DbHost").MustString("localhost")
		DbPort = Section.Key("DbPort").MustString("3306")
		DbUser = Section.Key("DbUser").MustString("nine")
		DbPassWord = Section.Key("DbPassWord").MustString("123qwe")
		DbName = Section.Key("DbName").MustString("ginblog")
	} else {
		var Section = file.Section(key_db)
		Db = Section.Key("Db").MustString("gee")
		DbHost = Section.Key("DbHost").MustString("localhost")
		DbPort = Section.Key("DbPort").MustString("3306")
		DbUser = Section.Key("DbUser").MustString("root")
		DbPassWord = Section.Key("DbPassWord").MustString("stan183183")
		DbName = Section.Key("DbName").MustString("ginblog")
	}

}
