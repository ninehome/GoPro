package model

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"ginblog/utils/errormsg"
	"io"

	//"github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
	Nickname string `gorm:"type:varchar(12)" json:"nickname" validate:"required,min=3,max=12" label:"昵称"`
	SEX      string `gorm:"type:varchar(6)" json:"sex" validate:"required,min=1,max=6" label:"昵称"`
	//Username string `gorm:"type:varchar(20);not null" json:"username"`
	//Password string `gorm:"type:varchar(20);not null" json:"password"`
	//Role     int    `gorm:"type:int" json:"role"` //用户角色

}

//{
//"username":"noe",
//"password":"123qwe",
//"role":2
//}

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

func LoginUser(name string, password string) (int, *User) {
	var user User
	password = strToMD5(password)
	code := CheckUser(name)

	if code == errormsg.ERROR_USERNAME_USED {
		fmt.Println("传入的用户名 和密码 -->", name, password)
		//查询 并且绑定到 users ,通过判断users的值 来判断是否查到记录
		//db.Where("username = ? ", name).First(&user)
		db.Where(&User{Username: name, Password: password}).First(&user)
		if user.ID > 0 {
			return errormsg.SUCCESS, &user //
		}
		return errormsg.ERROR_PASSWOR_WRONG, nil

	} else {
		return errormsg.ERROR_USER_NOT_EXIST, nil
	}

}

// 新增/注册用户 返回code
func CreateUser(data *User) (int, error) {
	data.Password = strToMD5(data.Password)
	//插入记录
	err := db.Create(&data).Error
	if err != nil {
		return errormsg.ERROR, err //500
	}
	return errormsg.SUCCESS, nil //200

}

//查询用户列表

func GetUserList(page int, pagesize int) []User {
	var ulist []User
	err := db.Limit(pagesize).Offset((page - 1) * pagesize).Find(&ulist).Error
	if err != nil || err == gorm.ErrRecordNotFound { //查询错误 和没记录 都返回nil
		return nil
	}

	return ulist
}

func EditeNickname(nickname string, username string) (int, error) {
	err := db.Model(&User{}).Where("username = ?", username).Update("nickname", nickname).Error
	if err != nil {
		return errormsg.ERROR, err
	}
	return errormsg.SUCCESS, nil
}

func strToMD5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
