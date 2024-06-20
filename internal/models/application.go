package models

import (
	"gorm.io/gorm"
	"time"
)

type Application struct {
	ID            string         `gorm:"column:id;type:varchar(128) not null primary key;comment:主键"`
	Name          string         `gorm:"column:name;type:varchar(64) not null;default:'';index:,unique;comment:名称"`
	Description   string         `gorm:"column:description;type:varchar(256) not null;default:'';comment:描述"`
	Secret        string         `gorm:"column:secret;type:varchar(256) not null;default:'';comment:秘钥"`
	Authorization bool           `gorm:"column:authorization;type:tinyint not null;default:1;comment:是否认证:0:不认证;1:认证"`
	CreatedAt     time.Time      `gorm:"column:created_at;type:timestamp not null;default:current_timestamp();comment:创建时间"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:timestamp not null default current_timestamp() on update current_timestamp();comment:最后一次更新时间"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;index;comment:是否删除"`
}

func (Application) TableName() string {
	return "application"
}
