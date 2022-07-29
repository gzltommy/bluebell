package logic

import (
	"bluebell/dao/mysql"
	"bluebell/model"
)

func GetCommunityList() ([]*model.Community, error) {
	return mysql.GetCommunityList()
}

func GetCommunityDetailByID(id uint64) (*model.CommunityDetail, error) {
	return mysql.GetCommunityByID(id)
}
