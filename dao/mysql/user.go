package mysql

import (
	"bluebell/models"
	"errors"
)

func CheckUserExist(username string) (error error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}

func InsertUser(user *models.User) (error error) {
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err := db.Exec(sqlStr, user.UserID, user.UserName, user.Password)
	return err
}
