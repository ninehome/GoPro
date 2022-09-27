package main

import (
	"GeeOrm/geeorm"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//fmt.Println("1111")
	//// DSN:Data Source Name
	//dsn := "user:password@tcp(127.0.0.1:3306)/dbname"
	//db, err := sql.Open("mysql", dsn)
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close() // 注意这行代码要写在上面err判断的下面

	engine, _ := geeorm.NewEngine("mysql", "root:stan183183@tcp(127.0.0.1:3306)/gee")
	//defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()    //如果存在 user表，删除
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec() //创建表
	//_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "toy", "nick").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
