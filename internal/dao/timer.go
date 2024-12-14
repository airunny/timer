package dao

import (
	"context"

	v1 "github.com/airunny/timer/api/timer/v1"
	"github.com/airunny/timer/internal/models"
	"github.com/airunny/wiki-go-tools/igorm"
	"github.com/airunny/wiki-go-tools/ormhelper"
	"gorm.io/gorm"
)

type TimerRecord struct {
	db *gorm.DB
}

func NewTimer(db *gorm.DB) *TimerRecord {
	return &TimerRecord{
		db: db,
	}
}

func (s *TimerRecord) Session(ctx context.Context, opts ...igorm.Option) *gorm.DB {
	return igorm.NewOptions(s.db, opts...).Session().WithContext(ctx)
}

func (s *TimerRecord) Add(ctx context.Context, in *models.Timer, opts ...igorm.Option) error {
	err := s.Session(ctx, opts...).
		Create(in).Error
	if err != nil {
		return ormhelper.WrapErr(err)
	}
	return nil
}

func (s *TimerRecord) Get(ctx context.Context, id string, opts ...igorm.Option) (*models.Timer, error) {
	var out models.Timer
	err := s.Session(ctx, opts...).
		Where("id = ?", id).
		First(&out).Error
	if err != nil {
		return nil, ormhelper.WrapErr(err)
	}
	return &out, nil
}

func (s *TimerRecord) Update(ctx context.Context, in *models.Timer, opts ...igorm.Option) error {
	err := s.Session(ctx, opts...).
		Updates(in).Error
	if err != nil {
		return ormhelper.WrapErr(err)
	}
	return nil
}

func (s *TimerRecord) UpdateStatus(ctx context.Context, id string, status v1.TimerStatus, opts ...igorm.Option) error {
	err := s.Session(ctx, opts...).
		Model(&models.Timer{}).
		Where("id = ?", id).
		Update("status", status).Error
	if err != nil {
		return ormhelper.WrapErr(err)
	}
	return nil
}

func (s *TimerRecord) Delete(ctx context.Context, id string, opts ...igorm.Option) error {
	err := s.Session(ctx, opts...).
		Delete(&models.Timer{
			ID: id,
		}).Error
	if err != nil {
		return ormhelper.WrapErr(err)
	}
	return nil
}

func (s *TimerRecord) Find(ctx context.Context, size int, offset string, opts ...igorm.Option) ([]*models.Timer, string, error) {
	var (
		out     []*models.Timer
		session = s.Session(ctx, opts...).
			Order("id desc")
		nextOffset = ""
	)

	if offset != "" {
		session = session.Where("id < ?", offset)
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
