package model

import "time"

// Post 内存对齐概念 字段类型相同的对齐 缩小变量所占内存大小
type Post struct {
	PostID      uint64    `json:"post_id,string" db:"post_id"`
	AuthorId    uint64    `json:"author_id" db:"author_id"`
	CommunityID uint64    `json:"community_id" db:"community_id"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title"`
	Content     string    `json:"content" db:"content"`
	CreateTime  time.Time `json:"-" db:"create_time"`
}
