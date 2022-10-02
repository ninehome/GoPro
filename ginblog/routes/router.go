package routes

import (
	"ginblog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")

	router.GET("hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})

	return r
}
