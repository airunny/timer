package models

import (
	"time"

	v1 "github.com/airunny/timer/api/timer/v1"
)

type Task struct {
	ID           int64         `gorm:"column:id;type:bigint unsigned not null primary key auto_increment;comment:主键"`
	TimerId      string        `gorm:"column:timer_id;type:varchar(128) not null default '';index:idx_timer_id_run_time;comment:定时器ID"`
	RunTime      time.Time     `gorm:"column:run_time;type:timestamp not null;default:current_timestamp();index:idx_timer_id_run_time;comment:执行时间"`
	Request      string        `gorm:"column:request;type:text;comment:回调内容"`
	Response     string        `gorm:"column:response;type:text;comment:回调返回内容"`
	Status       v1.TaskStatus `gorm:"column:status;type:smallint not null;default:0;index:idx_create_status,priority:2;comment:回调状态:0:成功;1:失败"`
	Retry        int64         `gorm:"column:retry;type:int not null default '0';comment:失败重试次数"`
	FailedReason string        `gorm:"column:failed_reason;type:text;comment:失败原因"`
	CreatedAt    time.Time     `gorm:"column:created_at;type:timestamp not null;default:current_timestamp();index:idx_create_status,priority:1;comment:创建时间"`
	UpdatedAt    time.Time     `gorm:"column:updated_at;type:timestamp not null;default:current_timestamp();comment:最后一次更新时间"`
}

func (Task) TableName() string {
	return "task"
}
