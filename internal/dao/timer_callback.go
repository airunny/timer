package dao

import (
	"context"
	"time"

	v1 "github.com/airunny/timer/api/timer/v1"
	"github.com/airunny/timer/internal/models"

	"github.com/airunny/wiki-go-tools/igorm"
	"github.com/airunny/wiki-go-tools/ormhelper"
	"gorm.io/gorm"
)

type TimerCallback struct {
	db *gorm.DB
}

func NewTimerCallback(db *gorm.DB) *TimerCallback {
	return &TimerCallback{
		db: db,
	}
}

func (s *TimerCallback) Session(ctx context.Context, opts ...igorm.Option) *gorm.DB {
	return igorm.NewOptions(s.db, opts...).Session().WithContext(ctx)
}

func (s *TimerCallback) Add(ctx context.Context, in *models.TimerCallback, opts ...igorm.Option) error {
	err := s.Session(ctx, opts...).Create(in).Error
	if err != nil {
		return ormhelper.WrapErr(err)
	}
	return nil
}

func (s *TimerCallback) FindByTimerId(ctx context.Context, timerId, offset string, size int, opts ...igorm.Option) ([]*models.TimerCallback, string, error) {
	var (
		out     []*models.TimerCallback
		session = s.Session(ctx, opts...).
			Where("timer_id = ?", timerId).
			Order("id desc")
		nextOffset = ""
	)

	if offset != "" {
		session = session.Where("timer_id < ?", offset)
	}

	err := session.Limit(size + 1).Find(&out).Error
	if err != nil {
		return out, nextOffset, ormhelper.WrapErr(err)
	}

	count := len(out)
	if count > 1 && count > size {
		nextOffset = out[count-2].ID
	}

	if count > size {
		out = out[:count-1]
	}
	return out, nextOffset, nil
}

func (s *TimerCallback) FindByTime(ctx context.Context, start, end time.Time, status v1.TimerCallbackStatus, opts ...igorm.Option) ([]*models.TimerCallback, error) {
	var out []*models.TimerCallback
	err := s.Session(ctx, opts...).
		Where("created_at >= ? and created_at <= ? and status = ?", start, end, status).
		Order("created_at desc").
		Find(&out).Error
	if err != nil {
		return nil, ormhelper.WrapErr(err)
	}
	return out, nil
}
