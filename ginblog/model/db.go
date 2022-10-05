package model

import (
	"fmt"
	"ginblog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
	"time"
)

var err error
var db *gorm.DB

// func InitDB() {
//
//	//"test:123qwe@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
//	sqlstr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
//		utils.DbUser,
//		utils.DbPassWord,
//		utils.DbHost,
//		utils.DbPort,
//		utils.Db)
//	fmt.Println(sqlstr)
//	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
//		utils.DbUser,
//		utils.DbPassWord,
//		utils.DbHost,
//		utils.DbPort,
//		utils.Db))
//	if err != nil {
//		fmt.Println("连接数据库失败:", err)
//	} else {
//		fmt.Println("链接数据库成功。。。")
//	}
//
//	db.SingularTable(true)                                       //建表 不加s ,比如模型是user - >users
//	db.AutoMigrate(&Aritcel{}, &User{}, &Category{}, &Comment{}) //关联模型
//	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
//	db.DB().SetMaxIdleConns(10)
//
//	// SetMaxOpenConns 设置打开数据库连接的最大数量。
//	db.DB().SetMaxOpenConns(100)
//
//	// SetConnMaxLifetime 设置了连接可复用的最大时间。
//	db.DB().SetConnMaxLifetime(10 * time.Second)
//
//	db.LogMode(true)
//	//defer db.Close()
//
// }
func InitDB() {

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		// gorm日志模式：silent
		Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
		os.Exit(1)
	}

	// 迁移数据表，在没有数据表结构变更时候，建议注释不执行
	_ = db.AutoMigrate(&User{}, &Category{}, Comment{})

	sqlDB, _ := db.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

}
