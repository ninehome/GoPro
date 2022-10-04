package errormsg

const (
	SUCCESS = 200
	ERROR   = 500
	//code ==1000   // 100开头 用户模块错误

	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWOR_WRONG    = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_NOT_EXIST  = 1004
	ERROR_TOKEN_OUT_TIME   = 1005
	ERROR_TOKNE_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007

	//cdoe ===2000  //2000开始文章模块错误
	// code ===3000 //3000开始分类模块错误
)

var codeMsg = map[int]string{
	SUCCESS:                "ok",
	ERROR:                  "fail",
	ERROR_USERNAME_USED:    "用户名已经存在",
	ERROR_PASSWOR_WRONG:    "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_NOT_EXIST:  "token 不存在",
	ERROR_TOKEN_OUT_TIME:   "token 已经过期",
	ERROR_TOKNE_WRONG:      "token 不正确",
	ERROR_TOKEN_TYPE_WRONG: "token 格式错误",
}

func GetErrorMsg(code int) string {
	return codeMsg[code]
}
