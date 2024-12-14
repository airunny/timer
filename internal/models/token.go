package models

import (
	"time"
)

type Token struct {
	ID               int64     `gorm:"column:id;type:bigint unsigned not null primary key auto_increment;comment:主键"`
	UserId           string    `gorm:"column:user_id;type:varchar(128) not null;default:'';index:;comment:用户ID"`
	AccessToken      string    `gorm:"column:access_token;type:varchar(512) not null;default:'';index:,unique;comment:用户token"`
	ExpiresIn        int64     `gorm:"column:expires_in;type:bigint unsigned not null;default:0;comment:token过期时间，单位秒"`
	RefreshToken     string    `gorm:"column:refresh_token;type:varchar(256) not null;default:'';index:,unique;comment:刷新token"`
	RefreshExpiresIn int64     `gorm:"column:refresh_expires_in;type:bigint unsigned not null;default:0;comment:刷新token过期时间，单位秒"`
	CreatedAt        time.Time `gorm:"column:created_at;type:timestamp not null;default:current_timestamp();comment:创建时间"`
	UpdatedAt        time.Time `gorm:"column:updated_at;type:timestamp not null default current_timestamp() on update current_timestamp();comment:最后一次更新时间"`
}

func (Token) TableName() string {
	return "token"
}
