package logic

import (
	"bluebell/model"
)

func CreatePost(p *model.ParamCreatePost) error {
	//// 1、 生成 post_id (生成帖子ID)
	//post := &model.Post{
	//	PostID:      uint64(snowflake.GenID()),
	//	AuthorId:    p.AuthorId,
	//	CommunityID: p.CommunityID,
	//	Status:      0,
	//	Title:       p.Title,
	//	Content:     p.Content,
	//	//CreateTime:  time.Time{},
	//}
	//
	//// 2、创建帖子保存到数据库
	//if err := mysql.CreatePost(post); err != nil {
	//	zap.L().Error("mysql.CreatePost(&post) failed", zap.Error(err))
	//	return err
	//}
	//community, err := mysql.GetCommunityNameByID(fmt.Sprint(post.CommunityID))
	//if err != nil {
	//	zap.L().Error("mysql.GetCommunityNameByID failed", zap.Error(err))
	//	return err
	//}
	// redis 存储帖子信息
	//if err := redis.CreatePost(
	//	post.PostID,
	//	post.AuthorId,
	//	post.Title,
	//	TruncateByWords(post.Content, 120),
	//	community.CommunityID); err != nil {
	//	zap.L().Error("redis.CreatePost failed", zap.Error(err))
	//	return err
	//}
	return nil
}
