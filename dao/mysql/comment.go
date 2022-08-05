package mysql

import (
	"bluebell/model"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func CreateComment(comment *model.Comment) (err error) {
	sqlStr := `insert into comment(
	comment_id, content, post_id, author_id, parent_id)
	values(?,?,?,?,?)`

	_, err = db.Exec(sqlStr, comment.CommentID, comment.Content, comment.PostID,
		comment.AuthorID, comment.ParentID)
	if err != nil {
		zap.L().Error("insert comment failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}

func GetCommentListByIDs(ids []string) (commentList []*model.Comment, err error) {
	sqlStr := `select comment_id, content, post_id, author_id, parent_id, create_time
	from comment
	where comment_id in (?)`

	// 动态填充 id
	query, args, err := sqlx.In(sqlStr, ids)
	if err != nil {
		return
	}
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用 Rebind() 重新绑定它
	query = db.Rebind(query)
	err = db.Select(&commentList, query, args...)
	return
}
