package dbmodels

type CommentRecord struct {
	BaseModel
	PostId     uint32 `gorm:"not null" json:"postId"`
	AuthorId   uint32 `gorm:"not null" json:"authorId"`
	AuthorName string `gorm:"not null;type:varchar(255)" json:"authorName"`
	Content    string `gorm:"not null;type:varchar(255)" json:"content"`
	FloorNum   uint32 `gorm:"not null" json:"floorNum"`
}
