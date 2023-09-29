package models

import "gorm.io/plugin/soft_delete"

type Audit struct {
	ID        uint64                `gorm:"column:id;primaryKey"`
	Message   string                `gorm:"column:message;type:text)"`
	CreatedAt int64                 `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt int64                 `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;default:null"`
}

func (Audit) TableName() string {
	return "audit"
}
