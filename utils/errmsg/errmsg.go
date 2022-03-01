package errmsg

const (
	SUCCSE = 200
	ERROR  = 500

	// code= 1000... 用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008
	// code = 4000... 购物车模块

	// code = 5000... 产品模块
	ERROR_PRODUCT_NOT_EXIST = 5001
	ERROR_PRICE_WRONG       = 5002

	//todo 其他模块错误开发时自行添加
	//
	ERROR_JSON_TYPE_WRONG = 9001
	ERROR_DATABASE_WRONG  = 9002
	ERROR_UPLOAD_WRONG    = 9003
)

var codeMsg = map[int]string{
	SUCCSE:                  "OK",
	ERROR:                   "FAIL",
	ERROR_USERNAME_USED:     "用户名已存在！",
	ERROR_PASSWORD_WRONG:    "密码错误",
	ERROR_USER_NOT_EXIST:    "用户不存在",
	ERROR_TOKEN_EXIST:       "TOKEN不存在,请重新登陆",
	ERROR_TOKEN_RUNTIME:     "TOKEN已过期,请重新登陆",
	ERROR_TOKEN_WRONG:       "TOKEN不正确,请重新登陆",
	ERROR_TOKEN_TYPE_WRONG:  "TOKEN格式错误,请重新登陆",
	ERROR_USER_NO_RIGHT:     "该用户无权限",
	ERROR_PRODUCT_NOT_EXIST: "产品不存在",
	ERROR_PRICE_WRONG:       "金额格式错误",
	ERROR_JSON_TYPE_WRONG:   "JSON格式错误",
	ERROR_DATABASE_WRONG:    "数据库错误，请重试",
	ERROR_UPLOAD_WRONG:      "文件上传错误，请重试",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
