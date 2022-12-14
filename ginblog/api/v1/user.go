package v1

import (
	"fmt"
	"ginblog/model"
	"ginblog/utils/errormsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 查询用户是否 存在
func UserExist(ctx *gin.Context) {

}

// 添加用户
func AddUser(ctx *gin.Context) {
	var data model.User
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		fmt.Println(err)
		panic(err)
		return
	}

	code := model.CheckUser(data.Username)
	if code == errormsg.SUCCESS {
		ecode, e := model.CreateUser(&data) //写入数据库
		if e != nil {
			code = ecode
			panic(e)
		}
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

//查询单个用户 /登录

func UserLogin(ctx *gin.Context) {
	var data *model.User
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		fmt.Println(err)
		return
	}

	code, data := model.LoginUser(data.Username, data.Password)

	//返回http json 数据
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errormsg.GetErrorMsg(code),
	})

}

// 查询用户列表
func GetUserList(ctx *gin.Context) {

	page, _ := strconv.Atoi(ctx.Query("page"))
	PageSize, _ := strconv.Atoi(ctx.Query("pagesize"))

	if PageSize == 0 {
		PageSize = -1 //当等于-1 就是取消 limit 条件，不需要分页
	}

	if page == 0 {
		page = 1
	}

	data, total := model.GetUserList(page, PageSize)
	code := errormsg.SUCCESS

	ctx.JSON(http.StatusOK, gin.H{

		"status":  code,
		"data":    data,
		"message": errormsg.GetErrorMsg(code),
		"total":   total,
	})
}

// 编辑用户
func EditUser(ctx *gin.Context) {
	nickname := ctx.Query("nickname")
	userName := ctx.Query("username")
	code := errormsg.ERROR_PARAMS
	if nickname == "" || userName == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errormsg.GetErrorMsg(code),
		})
		return
	}

	code, err := model.EditeNickname(nickname, userName)
	if code == errormsg.ERROR {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormsg.GetErrorMsg(code),
	})
}

//删除用户

func DeleteUser(ctx *gin.Context) {

}
