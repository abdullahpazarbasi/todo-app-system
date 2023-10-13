package driven_adapter_db

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	ID        string         `gorm:"column:id;primaryKey"`
	UserID    string         `gorm:"column:user_id;index"`
	Label     string         `gorm:"column:label;size:100"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
	Tags      []TodoTag      `gorm:"foreignKey:ID;references:ID"`
}
