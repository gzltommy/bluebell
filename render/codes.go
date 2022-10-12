package render

/*
* 10000～19999 区间表示参数错误

* 20000～29999 区间表示用户错误

* 30000～39999 区间表示接口异常


 */

const (
	CodeSucceed = 0

	CodeErrParams    = 10000
	CodeInvalidToken = 10001

	CodeUserExisted          = 20000
	CodeUserNotExit          = 20001
	CodePasswordWrong        = 20002
	CodeNotLogin             = 20003
	CodeInvalidCommunityID   = 20004
	CodeInvalidAuthorization = 20005

	CodeServerBusy  = 30000
	CodeServerError = 30001
)

var codeMsg = map[int]string{
	CodeSucceed: "Success",

	CodeErrParams:    "Error params",
	CodeInvalidToken: "Invalid token",

	CodeUserExisted:          "User existed",
	CodeUserNotExit:          "User not exit",
	CodePasswordWrong:        "Password wrong",
	CodeNotLogin:             "Not login",
	CodeInvalidCommunityID:   "Invalid community id",
	CodeInvalidAuthorization: "Invalid Authorization",

	CodeServerBusy:  "Server busy",
	CodeServerError: "Server error",
}
