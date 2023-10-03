package models

import "gorm.io/plugin/soft_delete"

type Hypervisors struct {
	ID        uint64                `gorm:"column:id;primaryKey"`
	Host      string                `gorm:"column:host;type:varchar(250);not null"`
	Port      uint16                `gorm:"column:port;not null"`
	Secret    string                `gorm:"column:secret;type:varchar(128);not null"`
	CreatedAt int64                 `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt int64                 `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;default:null;index:idx_deleted_at"`

	Capabilities []HypervisorCapabilities `gorm:"references:ID;foreignKey:HypervisorID"`
}

func (Hypervisors) TableName() string {
	return "hypervisors"
}

func (Hypervisors) Pick(capability string) string {
	return "test"
}
