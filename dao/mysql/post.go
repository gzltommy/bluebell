package mysql

import (
	"bluebell/model"
	"go.uber.org/zap"
)

// CreatePost 创建帖子
func CreatePost(post *model.Post) (err error) {
	sqlStr := `insert into post(post_id, title, content, author_id, community_id) values(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, post.PostID, post.Title, post.Content, post.AuthorId, post.CommunityID)
	if err != nil {
		zap.L().Error("insert post failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}
