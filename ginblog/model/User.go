package model

import (
	"ginblog/utils/errormsg"
	"github.com/jinzhu/gorm"
)

// 登录用户
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"` //用户角色

}

// 查询用户是否存在
func CheckUser(name string) int {
	var users User
	//查询 并且绑定到 users ,通过判断users的值 来判断是否查到记录
	db.Select("id").Where("Username = ? ", name).First(&users)
	if users.ID > 0 {
		return errormsg.ERROR_USERNAME_USED //已经存在 不可用
	}
	return errormsg.SUCCESS
}

// 新增/注册用户 返回code
func CreateUser(data *User) int {
	//插入记录
	err := db.Create(&data).Error
	if err != nil {
		return errormsg.ERROR //500
	}
	return errormsg.SUCCESS //200

}
