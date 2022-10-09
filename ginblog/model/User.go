package model

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"ginblog/utils/errormsg"
	"golang.org/x/crypto/bcrypt"
	"io"

	//"github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	//validate 数据验证
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"` //gte=2 大于等于 2
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
	fmt.Println("Create 1111111111111111")
	err := db.Create(&data).Error
	fmt.Println("Create 22222222222222222")
	if err != nil {
		return errormsg.ERROR, err //500
	}
	return errormsg.SUCCESS, nil //200

}

//gorm 钩子函数

func (u *User) BeforeCreate(*gorm.DB) error {
	fmt.Println("我是 db 调用create 之前会 回调的钩子函数")
	return nil
}

//查询用户列表

func GetUserList(page int, pagesize int) ([]User, int) {
	var ulist []User
	var total int64
	err := db.Limit(pagesize).Offset((page - 1) * pagesize).Find(&ulist).Count(&total).Error
	if err != nil || err == gorm.ErrRecordNotFound { //查询错误 和没记录 都返回nil
		return nil, 0
	}

	return ulist, int(total)
}

func EditeNickname(nickname string, username string) (int, error) {
	err := db.Model(&User{}).Where("username = ?", username).Update("nickname", nickname).Error
	if err != nil {
		return errormsg.ERROR, err
	}
	return errormsg.SUCCESS, nil
}

// 密码加密方式1
func strToMD5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

//密码加密方式2  bcrypt

// 需要将第一次生成的hash 存入数据库，用来比较 之后的输入值
func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {

	}
	return string(hash)
}

// 验证
func ValidatePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}

//密码加密方式3 scrypt
