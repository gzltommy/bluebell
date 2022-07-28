package mysql

import (
	"bluebell/models"
	"database/sql"
)

func CheckUserExist(username string) (error error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExited
	}
	return
}

func InsertUser(user *models.User) (error error) {
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err := db.Exec(sqlStr, user.UserID, user.UserName, user.Password)
	return err
}

func Login(username, password string) (user *models.User, err error) {
	sqlStr := "select user_id, username, password from user where username = ?"
	err = db.Get(user, sqlStr, username)
	if err != nil {
		if err == sql.ErrNoRows {
			// 用户不存在
			return nil, ErrorUserNotExit
		} else {
			return nil, err
		}
	}
	// 生成加密密码与查询到的密码比较
	if user.Password != password {
		return nil, ErrorPasswordWrong
	}
	return
}
