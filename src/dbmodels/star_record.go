package dbmodels

import "time"

type StarRecord struct {
	ID        uint32    `gorm:"primaryKey;not null" json:"id"`
	PostId    uint32    `gorm:"not null;index:idx_star" json:"postId"`
	UserId    uint32    `gorm:"not null;index:idx_star" json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
}
