package dao

import (
	"context"
	"strconv"
	"time"

	v1 "github.com/airunny/timer/api/timer/v1"
	innErr "github.com/airunny/timer/errors"
	"github.com/airunny/timer/internal/models"
	"github.com/airunny/wiki-go-tools/igorm"
	"github.com/airunny/wiki-go-tools/ormhelper"
	"gorm.io/gorm"
)

type Task struct {
	db *gorm.DB
}

func NewTask(db *gorm.DB) *Task {
	return &Task{
		db: db,
	}
}

func (s *Task) session(ctx context.Context, opts ...igorm.Option) *gorm.DB {
	return igorm.NewOptions(s.db, opts...).Session().WithContext(ctx)
}

func (s *Task) Add(ctx context.Context, in *models.Task, opts ...igorm.Option) (int64, error) {
	err := s.session(ctx, opts...).Debug().Create(in).Error
	if err != nil {
		return 0, ormhelper.WrapErr(err)
	}
	return in.ID, nil
}

func (s *Task) FindByTimerId(ctx context.Context, timerId, offset string, size int, status v1.TaskStatus, opts ...igorm.Option) ([]*models.Task, string, error) {
	var (
		out        []*models.Task
		nextOffset string
		session    = s.session(ctx, opts...).
				Where("timer_id = ?", timerId).
				Order("id desc")
	)

	if status > 0 {
		session = session.Where("status = ?", status)
	}

	if offset != "" {
		intOffset, err := strconv.ParseInt(offset, 10, 64)
		if err != nil {
			return nil, "", innErr.ErrBadRequest
		}
		session = session.Where("id < ?", intOffset)
	}

	err := session.Limit(size + 1).Find(&out).Error
	if err != nil {
		return out, nextOffset, ormhelper.WrapErr(err)
	}

	count := len(out)
	if count > 1 && count > size {
		nextOffset = strconv.Itoa(int(out[count-2].ID))
	}

	if count > size {
		out = out[:count-1]
	}
	return out, nextOffset, nil
}

func (s *Task) FindByTime(ctx context.Context, offset string, start, end time.Time, size int, status v1.TaskStatus, opts ...igorm.Option) ([]*models.Task, string, error) {
	var (
		out        []*models.Task
		nextOffset string
		session    = s.session(ctx, opts...).
				Where("created_at >= ? and created_at <= ?", start, end).
				Order("id desc")
	)

	if status > 0 {
		session = session.Where("status = ?", status)
	}

	if offset != "" {
		intOffset, err := strconv.ParseInt(offset, 10, 64)
		if err != nil {
			return nil, "", innErr.ErrBadRequest
		}
		session = session.Where("id < ?", intOffset)
	}

	err := session.Limit(size + 1).Debug().Find(&out).Error
	if err != nil {
		return out, nextOffset, ormhelper.WrapErr(err)
	}

	count := len(out)
	if count > 1 && count > size {
		nextOffset = strconv.Itoa(int(out[count-2].ID))
	}

	if count > size {
		out = out[:count-1]
	}
	return out, nextOffset, nil
}
