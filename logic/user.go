package logic

import (
	"bluebell/dao/mysql"
	"bluebell/model"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
	"bluebell/utils"
)

func SignUp(p *model.ParamSignUp) error {
	// 1. 查看用户是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}

	// 2.生成 UID
	uid := snowflake.GenID()

	// 构造一个 User 实例
	user := model.User{
		UserID:   uint64(uid),
		UserName: p.Username,
		Password: utils.EncryptPassword([]byte(p.Password)), // 对密码进行加密
	}

	// 3. 插入数据库
	return mysql.InsertUser(&user)
}

func Login(p *model.ParamLogin) (user *model.User, err error) {
	user, err = mysql.Login(p.UserName, utils.EncryptPassword([]byte(p.Password)))
	if err != nil {
		return
	}

	// 生成 JWT
	//return jwt.GenToken(user.UserID,user.UserName)
	aToken, rToken, err := jwt.GenToken(user.UserID, user.UserName)
	user.AccessToken = aToken
	user.RefreshToken = rToken
	return

}
