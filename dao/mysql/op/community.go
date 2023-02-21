package op

import (
	"bluebell/dao/mysql"
	"bluebell/model"
	"database/sql"
	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*model.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	err = mysql.DB.Select(&communityList, sqlStr)
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
	err = mysql.DB.Get(community, sqlStr, id)
	if err == sql.ErrNoRows { // 查询为空
		err = mysql.ErrorInvalidID // 无效的ID
		return
	}
	if err != nil {
		zap.L().Error("query community failed", zap.String("sql", sqlStr), zap.Error(err))
		err = mysql.ErrorQueryFailed
	}
	return community, err
}

func GetCommunityNameByID(idStr string) (community *model.Community, err error) {
	community = new(model.Community)
	sqlStr := `select community_id, community_name	from community	where community_id = ?`
	err = mysql.DB.Get(community, sqlStr, idStr)
	if err == sql.ErrNoRows {
		err = mysql.ErrorInvalidID
		return
	}
	if err != nil {
		zap.L().Error("query community failed", zap.String("sql", sqlStr), zap.Error(err))
		err = mysql.ErrorQueryFailed
		return
	}
	return
}
