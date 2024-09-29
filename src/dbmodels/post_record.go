package dbmodels

import "time"

type PostRecord struct {
	BaseModel
	AuthorId        uint32    `gorm:"not null" json:"authorId"`
	AuthorName      string    `gorm:"not null;type:varchar(255)" json:"authorName"`
	CommentCount    uint32    `gorm:"not null" json:"commentCount"`
	UpvoteCount     uint32    `gorm:"not null" json:"upvoteCount"`
	Title           string    `gorm:"not null;type:varchar(255)" json:"title"`
	ContentJson     string    `gorm:"not null;type:json" json:"contentJson"`
	LatestRepliedAt time.Time `json:"latestRepliedAt"`
}
