package render

/*
* -1 未定义的错误

* 10000～19999 区间表示参数错误

* 20000～29999 区间表示用户错误

* 30000～39999 区间表示接口异常


 */

const (
	ErrorUndefined = -1
	Succeed        = 0

	ErrParams = 10000

	UserExisted = 20000
)

var codeMsg = map[int]string{
	ErrorUndefined: "Undefined error",
	Succeed:        "Success",

	ErrParams:   "Error params",
	UserExisted: "User existed",
}
