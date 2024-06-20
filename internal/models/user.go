package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         int64          `json:"id" gorm:"column:Id"`                                 // id
	CreateTime time.Time      `json:"create_time" gorm:"autoCreateTime;column:CreateTime"` // 创建时间
	UpdateTime time.Time      `json:"update_time" gorm:"autoUpdateTime;column:UpdateTime"` // 更新时间
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"column:DeletedAt"`                  // 是否删除
}
