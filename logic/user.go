package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"bluebell/utils"
	"errors"
)

var (
	ErrorUserExited = errors.New("user exited")
)

func SignUp(p *models.ParamSignUp) error {
	// 1. 查看用户是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}

	// 2.生成 UID
	uid := snowflake.GenID()

	// 构造一个 User 实例
	user := models.User{
		UserID:   uint64(uid),
		UserName: p.Username,
		Password: utils.EncryptPassword([]byte(p.Password)), // 对密码进行加密
	}

	// 3. 插入数据库
	return mysql.InsertUser(&user)
}
