package dbmodels

import (
	"gorm.io/gorm"
	"time"
)

// BaseModel 软删除支持
type BaseModel struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
