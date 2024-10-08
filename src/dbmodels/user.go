package dbmodels

type User struct {
	BaseModel
	Username string `gorm:"unique;not null"` // 用户名，唯一且不能为空
	Password string // 密码
	Email    string `gorm:"unique"` // 邮箱，唯一
	Avatar   string // 用户头像路径或 URL
	Bio      string // 用户简介
}
