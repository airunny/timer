package dao

import (
	"context"
	"github.com/airunny/wiki-go-tools/igorm"
	"github.com/airunny/wiki-go-tools/ormhelper"
	"gorm.io/gorm"
	"timer/internal/models"
)

type TimerRecord struct {
	db *gorm.DB
}

func NewTimerRecord(db *gorm.DB) *TimerRecord {
	return &TimerRecord{
		db: db,
	}
}

func (s *TimerRecord) Session(opts ...igorm.Option) *gorm.DB {
	return igorm.NewOptions(s.db, opts...).Session()
}

func (s *TimerRecord) Add(ctx context.Context, in *models.TimerRecord, opts ...igorm.Option) error {
	err := s.Session(opts...).WithContext(ctx).Create(in).Error
	if err != nil {
		return ormhelper.WrapErr(err)
	}
	return nil
}

func (s *TimerRecord) Get(ctx context.Context, id string, opts ...igorm.Option) (*models.TimerRecord, error) {
	var out models.TimerRecord
	err := s.Session(opts...).WithContext(ctx).Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, ormhelper.WrapErr(err)
	}
	return &out, nil
}

func (s *TimerRecord) Update(ctx context.Context, in *models.TimerRecord, opts ...igorm.Option) error {
	err := s.Session(opts...).WithContext(ctx).Updates(in).Error
	if err != nil {
		return ormhelper.WrapErr(err)
	}
	return nil
}

func (s *TimerRecord) Delete(ctx context.Context, id string, opts ...igorm.Option) error {
	err := s.Session(opts...).WithContext(ctx).Delete(&models.TimerRecord{
		ID: id,
	}).Error
	if err != nil {
		return ormhelper.WrapErr(err)
	}
	return nil
}

func (s *TimerRecord) Find(ctx context.Context, size int, offset string, opts ...igorm.Option) ([]*models.TimerRecord, string, error) {
	var (
		out        []*models.TimerRecord
		session    = s.Session(opts...).WithContext(ctx).Order("id desc")
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
