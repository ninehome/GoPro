package model

import (
	"fmt"
	"ginblog/utils/errormsg"
	//"github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

// 登录用户
type User struct {
	gorm.Model

	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
	//Username string `gorm:"type:varchar(20);not null" json:"username"`
	//Password string `gorm:"type:varchar(20);not null" json:"password"`
	//Role     int    `gorm:"type:int" json:"role"` //用户角色

}

// {
// "username":"noe",
// "password":"123qwe",
// "role":2
// }

// 查询用户是否存在
func CheckUser(name string) int {
	var user User
	fmt.Println("传入的用户名是 -->", name)
	//查询 并且绑定到 users ,通过判断users的值 来判断是否查到记录
	//db.Where("username = ? ", name).First(&user)
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
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
