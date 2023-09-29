package models

import "gorm.io/plugin/soft_delete"

type Users struct {
	ID        uint64                `gorm:"column:id;primaryKey"`
	Email     string                `gorm:"column:email;type:varchar(250);index:idx_email,unique;not null"`
	Password  string                `gorm:"column:password;type:varchar(60);not null"`
	CreatedAt int64                 `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt int64                 `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;default:null;index:idx_deleted_at"`
}

func (Users) TableName() string {
	return "users"
}
