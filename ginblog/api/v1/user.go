package v1

import (
	"ginblog/model"
	"ginblog/utils/errormsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 查询用户是否 存在
func UserExist(ctx *gin.Context) {

}

// 添加用户
func AddUser(ctx *gin.Context) {
	var data model.User
	_ = ctx.ShouldBindJSON(&data)
	code := model.CheckUser(data.Username)
	if code == errormsg.SUCCESS {
		model.CreateUser(&data) //写入数据库
	}

	if code == errormsg.ERROR_USERNAME_USED {
		code = errormsg.ERROR_USERNAME_USED
	}

	//返回http json 数据
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errormsg.GetErrorMsg(code),
	})
}

//查询单个用户

// 查询用户列表
func GetUserList(ctx *gin.Context) {

}

// 编辑用户
func EditUser(ctx *gin.Context) {

}

//删除用户

func DeleteUser(ctx *gin.Context) {

}
