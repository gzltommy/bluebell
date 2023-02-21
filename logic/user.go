package logic

import (
	model2 "bluebell/dao/mysql/op"
	"bluebell/model"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
	"bluebell/utils"
)

func SignUp(p *model.ParamSignUp) error {
	// 1. 查看用户是否存在
	if err := model2.CheckUserExist(p.Username); err != nil {
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
	return model2.InsertUser(&user)
}

func Login(p *model.ParamLogin) (user *model.User, err error) {
	user, err = model2.Login(p.UserName, utils.EncryptPassword([]byte(p.Password)))
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
