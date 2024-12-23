package models

import (
	"time"

	v1 "github.com/airunny/timer/api/timer/v1"
)

type User struct {
	ID             string        `gorm:"column:id;type:varchar(128) not null primary key;comment:主键"`
	Name           string        `gorm:"column:name;type:varchar(64) not null;default:'';index:,unique;comment:名称"`
	Password       string        `gorm:"column:password;type:varchar(256) not null;default:'';comment:密码"`
	Status         v1.UserStatus `gorm:"column:status;type:smallint not null;default:0;comment:应用状态:0:未知;1:禁用;2:启用"`
	Role           v1.UserRole   `gorm:"column:role;type:smallint not null;default:0;comment:用户角色:0:未知;1:管理员;2:读写;3:只读"`
	ChangePassword bool          `gorm:"column:change_password;type:smallint not null;default:0;comment:是否更改密码"`
	CreatedAt      time.Time     `gorm:"column:created_at;type:timestamp not null;default:current_timestamp();comment:创建时间"`
	UpdatedAt      time.Time     `gorm:"column:updated_at;type:timestamp not null default current_timestamp() on update current_timestamp();comment:最后一次更新时间"`
}

func (User) TableName() string {
	return "user"
}
