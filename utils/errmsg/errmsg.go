package errmsg

const (
	SUCCESS = 200
	ERROR = 500

	// code = 1000... 用户模块的错误
	ErrorUsernameUsed   = 1001
	ErrorPasswordWrong  = 1002
	ErrorUserNotExist   = 1003
	ErrorTokenNotExist     = 1004
	ErrorTokenRuntime   = 1005
	ErrorTokenWrong     = 1006
	ErrorTokenTypeWrong = 1007
	ErrorUserNoRight = 1008
	ErrorUserOrPasswordValidateWrong = 1009

	// code = 2000... 文章模块的错误
	ErrorArtNotExist = 2001
	// code = 3000... 分类模块的错误
	ErrorCategoryUsed = 3001
	ErrorCategoryNotExist = 3002
	// code = 4000... 学生模块
	ErrorStudentNameUsed = 4001
	ErrorStudentNameNotExist = 4002
	// code = 5000... 书籍模块
	ErrorBookNameUsed = 5001
)

var codeMsg = map[int]string{
	SUCCESS:             "OK",
	ERROR:               "FAILED",
	ErrorUsernameUsed:   "用户名已存在",
	ErrorPasswordWrong:  "密码错误",
	ErrorUserNotExist:   "用户不存在",
	ErrorTokenNotExist:     "TOKEN不存在",
	ErrorTokenRuntime:   "TOKEN已过期",
	ErrorTokenWrong:     "TOKEN不正确",
	ErrorTokenTypeWrong: "TOKEN格式错误",
	ErrorCategoryUsed: "分类已存在",
	ErrorArtNotExist: "文章不存在",
	ErrorCategoryNotExist: "该分类不存在",
	ErrorUserNoRight: "用户没有权限",
	ErrorUserOrPasswordValidateWrong: "用户名或者密码格式错误",
	ErrorStudentNameUsed: "学生名已用",
	ErrorStudentNameNotExist: "学生外键不存在",
	ErrorBookNameUsed: "书名已用",
}

func GetErrMsg(code int) string  {
	return codeMsg[code]
}