package mysql

import (
	"bluebell/model"
	"database/sql"
	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*model.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	err = db.Select(&communityList, sqlStr)
	if err == sql.ErrNoRows { // 查询为空
		zap.L().Warn("there is no community in db")
		err = nil
	}
	return
}

func GetCommunityByID(id uint64) (community *model.CommunityDetail, err error) {
	community = new(model.CommunityDetail)
	sqlStr := `select community_id, community_name, introduction, create_time
					from community
				where community_id = ?`
	err = db.Get(community, sqlStr, id)
	if err == sql.ErrNoRows { // 查询为空
		err = ErrorInvalidID // 无效的ID
		return
	}
	if err != nil {
		zap.L().Error("query community failed", zap.String("sql", sqlStr), zap.Error(err))
		err = ErrorQueryFailed
	}
	return community, err
}

func GetCommunityNameByID(idStr string) (community *model.Community, err error) {
	community = new(model.Community)
	sqlStr := `select community_id, community_name	from community	where community_id = ?`
	err = db.Get(community, sqlStr, idStr)
	if err == sql.ErrNoRows {
		err = ErrorInvalidID
		return
	}
	if err != nil {
		zap.L().Error("query community failed", zap.String("sql", sqlStr), zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	return
}
