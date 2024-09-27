package models

import (
	"time"

	v1 "github.com/airunny/timer/api/timer/v1"
)

type TimerCallback struct {
	ID             string                 `gorm:"column:id;type:varchar(128) not null primary key;comment:主键"`
	TimerId        string                 `gorm:"column:timer_id;type:varchar(128) not null default '';index;comment:定时器ID"`
	RequestContent string                 `gorm:"column:request_content;type:text;comment:回调内容"`
	ReplyContent   string                 `gorm:"column:reply_content;type:text;comment:回调返回内容"`
	FailedReason   string                 `gorm:"column:failed_reason;type:text;comment:失败原因"`
	RetryCount     int64                  `gorm:"column:retry_count;type:int not null default '0';comment:失败重试次数"`
	Status         v1.TimerCallbackStatus `gorm:"column:status;type:smallint not null;default:0;index:idx_create_status,priority:2;comment:回调状态:0:成功;1:失败"`
	CreatedAt      time.Time              `gorm:"column:created_at;type:timestamp not null;default:current_timestamp();index:idx_create_status,priority:1;comment:创建时间"`
}

func (TimerCallback) TableName() string {
	return "timer_callback"
}
