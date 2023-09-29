package models

import "gorm.io/plugin/soft_delete"

type Users struct {
	ID        uint64                `gorm:"column:id;primaryKey"`
	Email     string                `gorm:"column:email;index:idx_email,unique;not null"`
	Password  string                `gorm:"column:password;not null"`
	JWTSecret string                `gorm:"column:jwt_secret;not null"`
	CreatedAt int64                 `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt int64                 `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;default:null"`
}

func (Users) TableName() string {
	return "users"
}
