package model

const (
	OrderTime  = "time"
	OrderScore = "score"
)

type (
	ParamSignUp struct {
		Username        string `json:"username" binding:"required"`
		Password        string `json:"password" binding:"required"`
		ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
	}

	ParamLogin struct {
		UserName string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	ParamCreatePost struct {
		AuthorId    uint64 `json:"author_id"`
		CommunityID uint64 `json:"community_id" binding:"required"`
		Title       string `json:"title" binding:"required"`
		Content     string `json:"content" binding:"required"`
	}

	ParamPostList struct {
		Page int64 `json:"page" form:"page"` // 页码
		Size int64 `json:"size" form:"size"` // 每页数量
	}

	ParamPostList2 struct {
		CommunityID uint64 `json:"community_id" form:"community_id"`   // 可以为空
		Page        int64  `json:"page" form:"page"`                   // 页码
		Size        int64  `json:"size" form:"size"`                   // 每页数量
		Order       string `json:"order" form:"order" example:"score"` // 排序依据
	}

	ParamVote struct {
		PostID    string `json:"post_id" binding:"required"`              // 帖子id
		Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` // 赞成票(1)还是反对票(-1)取消投票(0)
	}
)
