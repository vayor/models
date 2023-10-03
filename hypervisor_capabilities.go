package models

import "gorm.io/plugin/soft_delete"

type HypervisorCapabilities struct {
	ID           uint64                `gorm:"column:id;primaryKey"`
	HypervisorID uint64                `gorm:"column:hypervisor_id;default:null"`
	Capability   string                `gorm:"column:capability;type:enum('container');default:null"`
	CreatedAt    int64                 `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    int64                 `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt    soft_delete.DeletedAt `gorm:"column:deleted_at;default:null;index:idx_deleted_at"`
}

func (HypervisorCapabilities) TableName() string {
	return "hypervisor_capabilities"
}
