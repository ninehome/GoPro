package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	//r := gin.New()
	r.Use(gin.Recovery())
	router := r.Group("api/v1")
	{
		//user模块接口
		router.POST("user/add", v1.AddUser)
		router.POST("user/login", v1.UserLogin)
		router.GET("user/users", v1.GetUserList)
		router.PUT("user/:id", v1.EditUser)
		router.GET("user/nickname", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)

		//分类模块接口

		//文章模块路由接口
	}

	//router.GET("hello", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"msg": "ok",
	//	})
	//})

	return r
}
