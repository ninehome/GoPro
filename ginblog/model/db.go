package model

import (
	"fmt"
	"ginblog/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var err error
var db *gorm.DB

func InitDB() {

	//"test:123qwe@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
	sqlstr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.Db)
	fmt.Println(sqlstr)
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.Db))
	if err != nil {
		fmt.Println("连接数据库失败:", err)
	} else {
		fmt.Println("链接数据库成功。。。")
	}

	db.SingularTable(true)                           //建表 不加s ,比如模型是user - >users
	db.AutoMigrate(&Aritcel{}, &User{}, &Category{}) //关联模型
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	db.DB().SetConnMaxLifetime(10 * time.Second)

	db.LogMode(true)
	//defer db.Close()

}
