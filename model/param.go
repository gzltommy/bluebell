package model

type ParamSignUp struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamCreatePost struct {
	AuthorId    uint64 `json:"author_id"`
	CommunityID uint64 `json:"community_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
}
