package domain

import (
	"time"

	"gorm.io/gorm"
)

type (
	Base struct {
		ID        string         `gorm:"column:id;type:varchar(64);primaryKey"`
		CreatedAt time.Time      `gorm:"<-:create;column:created_at;autoCreate"`
		UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdate"`
		DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
	}
)
