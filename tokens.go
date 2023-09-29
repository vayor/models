package models

import "gorm.io/plugin/soft_delete"

type Tokens struct {
	ID        uint64                `gorm:"column:id;primaryKey"`
	UserID    uint64                `gorm:"column:user_id"`
	Scope     string                `gorm:"column:scope;type:enum('internal','external');default:internal"`
	Secret    string                `gorm:"column:secret;type:varchar(36)"`
	CreatedAt int64                 `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt int64                 `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;default:null"`
}

func (Tokens) TableName() string {
	return "tokens"
}
