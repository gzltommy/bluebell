package op

import (
	"bluebell/dao/mysql"
	"bluebell/model"
	"database/sql"
)

func CheckUserExist(username string) (error error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := mysql.DB.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return mysql.ErrorUserExited
	}
	return
}

func InsertUser(user *model.User) (error error) {
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err := mysql.DB.Exec(sqlStr, user.UserID, user.UserName, user.Password)
	return err
}

func Login(username, password string) (user *model.User, err error) {
	user = new(model.User)
	sqlStr := "select user_id, username, password from user where username = ?"
	err = mysql.DB.Get(user, sqlStr, username)
	if err != nil {
		if err == sql.ErrNoRows {
			// 用户不存在
			return nil, mysql.ErrorUserNotExit
		} else {
			return nil, err
		}
	}
	// 生成加密密码与查询到的密码比较
	if user.Password != password {
		return nil, mysql.ErrorPasswordWrong
	}
	return
}

func GetUserByID(id uint64) (user *model.User, err error) {
	user = new(model.User)
	sqlStr := `select user_id, username from user where user_id = ?`
	err = mysql.DB.Get(user, sqlStr, id)
	return
}
