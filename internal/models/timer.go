package models

import (
	"database/sql/driver"
	"time"

	v1 "github.com/airunny/timer/api/timer/v1"

	"github.com/airunny/wiki-go-tools/igorm"
	"gorm.io/gorm"
)

type Timer struct {
	ID              string          `gorm:"column:id;type:varchar(128) not null primary key;comment:主键"`
	AppId           string          `gorm:"column:app_id;type:varchar(64) not null;default:'';index:;comment:应用ID"`
	ExpireAt        int64           `gorm:"column:expire_at;type:bigint not null;default:0;comment:定时器时间"`
	Attach          string          `gorm:"column:attach;type:text;comment:附件内容"`
	CallbackType    v1.CallbackType `gorm:"column:callback_type;type:tinyint not null;default:0;comment:回调类型:0:http;1:mq"`
	CallbackAddress string          `gorm:"column:callback_address;type:varchar(128) not null;default:'';comment:回调地址"`
	Status          v1.TimerStatus  `gorm:"column:status;type:tinyint not null;default:0;comment:状态:0:未到时间;1:回调成功;2:失败"`
	RetryCount      int64           `gorm:"column:retry_count;type:int not null;default:0;comment:失败重试次数"`
	FailReason      string          `gorm:"column:fail_reason;type:text;comment:json格式失败原因"`
	Extra           *TimerExtra     `gorm:"column:extra;type:text;comment:json格式扩展字段"`
	CreatedAt       time.Time       `gorm:"column:created_at;comment:创建时间"`
	UpdatedAt       time.Time       `gorm:"column:updated_at;comment:最后一次更新时间"`
	DeletedAt       gorm.DeletedAt  `gorm:"column:deleted_at;index;comment:是否删除"`
}

func (Timer) TableName() string {
	return "timer"
}

type TimerExtra struct {
	Params []*v1.Params `json:"params"`
}

func (j *TimerExtra) Value() (driver.Value, error) {
	return igorm.GormCustomValue(j)
}

func (j *TimerExtra) Scan(value interface{}) error {
	return igorm.GormCustomScan(j, value)
}