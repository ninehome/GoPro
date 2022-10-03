package routes

import (
	"ginblog/utils"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func InitRouter() *gin.Engine {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	// 记录到文件。
	f, _ := os.Create("gin.log")
	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := r.Group("api/v1")

	router.GET("hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})

	return r
}
