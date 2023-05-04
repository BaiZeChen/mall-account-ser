package data

import "gorm.io/plugin/soft_delete"

type Base struct {
	ID         uint                  `gorm:"primaryKey;column:id"`
	CreateTime int                   `gorm:"autoCreateTime;column:create_time"`
	UpdateTime int                   `gorm:"autoUpdateTime;column:update_time"`
	IsDelete   soft_delete.DeletedAt `gorm:"softDelete:flag;column:is_delete"`
}
