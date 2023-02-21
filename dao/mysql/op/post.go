package op

import (
	"bluebell/dao/mysql"
	"bluebell/model"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"strings"
)

func GetPostList(page, size int) (posts []*model.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	ORDER BY create_time
	DESC 
	limit ?,?
	`
	posts = make([]*model.Post, 0, 2) // 0：长度  2：容量
	err = mysql.DB.Select(&posts, sqlStr, (page-1)*size, size)
	return

}

// CreatePost 创建帖子
func CreatePost(post *model.Post) (err error) {
	sqlStr := `insert into post(post_id, title, content, author_id, community_id) values(?,?,?,?,?)`

	_, err = mysql.DB.Exec(sqlStr, post.PostID, post.Title, post.Content, post.AuthorId, post.CommunityID)
	if err != nil {
		zap.L().Error("insert post failed", zap.Error(err))
		err = mysql.ErrorInsertFailed
		return
	}
	return
}

func GetPostByID(pid int64) (post *model.Post, err error) {
	post = new(model.Post)
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	where post_id = ?`

	err = mysql.DB.Get(post, sqlStr, pid)
	if err == sql.ErrNoRows {
		err = mysql.ErrorInvalidID
		return
	}
	if err != nil {
		zap.L().Error("query post failed", zap.String("sql", sqlStr), zap.Error(err))
		err = mysql.ErrorQueryFailed
		return
	}
	return
}

func GetPostListByIDs(ids []string) (postList []*model.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	where post_id in (?)
	order by FIND_IN_SET(post_id, ?)`

	// 动态填充id
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return
	}
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = mysql.DB.Rebind(query)
	err = mysql.DB.Select(&postList, query, args...)
	return
}
