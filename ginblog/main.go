package main

import (
	"fmt"
	"ginblog/model"
	"ginblog/routes"
	"ginblog/utils"
)

func main() {

	//RewriteStderrFile()
	//testPanic()

	//观察者模式
	//shirtItem := obsiver.NewItem("Basketball")
	//observerFirst := &obsiver.Customer{Id: "abc@gmail.com", Name: "李文"}
	//observerSecond := &obsiver.Customer{Id: "xyz@gmail.com", Name: "刘小"}
	//shirtItem.Register(observerFirst)
	//shirtItem.Register(observerSecond)
	//shirtItem.UpdateAvailability()

	divide()
	fmt.Println("divide 方法调用完毕，回到 main 函数")

	//web 服务  	//引用数据库
	model.InitDB()
	engin := routes.InitRouter()
	engin.Run(utils.HttpPort)

}

func divide() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Runtime panic caught: %v\n", err)
		}
	}()
	var i = 1
	var j = 0
	k := i / j
	fmt.Printf("%d / %d = %d\n", i, j, k)
}
