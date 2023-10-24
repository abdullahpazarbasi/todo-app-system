package driven_adapter_db

import (
	"gorm.io/gorm"
	"time"
)

type TodoTag struct {
	ID        string         `gorm:"column:id;primaryKey"`
	TodoID    string         `gorm:"column:todo_id;index"`
	Key       string         `gorm:"column:key;size:32"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
	Todo      Todo           `gorm:"references:ID;foreignKey:TodoID"`
}

func (TodoTag) TableName() string {
	return "todo_tag"
}
