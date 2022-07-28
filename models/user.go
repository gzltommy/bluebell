package models

type User struct {
	UserID       uint64 `json:"user_id,string" db:"user_id"` // 指定json序列化/反序列化时使用小写user_id
	UserName     string `json:"username" db:"username"`
	Password     string `json:"password" db:"password"`
	AccessToken  string `json:"access_token" db:"-"`
	RefreshToken string `json:"refresh_token" db:"-"`
}
