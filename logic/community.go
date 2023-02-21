package logic

import (
	model2 "bluebell/dao/mysql/op"
	"bluebell/model"
)

func GetCommunityList() ([]*model.Community, error) {
	return model2.GetCommunityList()
}

func GetCommunityDetailByID(id uint64) (*model.CommunityDetail, error) {
	return model2.GetCommunityByID(id)
}
