package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(utils.AppMode)

	/**func Default() *Engine {
		debugPrintWARNINGDefault()
		engine := New()
		engine.Use(Logger(), Recovery())
		return engine
	}

	*/
	//r := gin.Default()  //这个初始化包含了 gin 默认日志中间件
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.LoggerToFile())
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
