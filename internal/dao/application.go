package dao

import (
	"context"

	"github.com/airunny/wiki-go-tools/igorm"
	"github.com/airunny/wiki-go-tools/ormhelper"
	"gorm.io/gorm"
	"timer/internal/models"
)

type Application struct {
	db *gorm.DB
}

func NewApplication(db *gorm.DB) *Application {
	return &Application{
		db: db,
	}
}

func (s *Application) Session(opts ...igorm.Option) *gorm.DB {
	return igorm.NewOptions(s.db, opts...).Session()
}

func (s *Application) Add(ctx context.Context, in *models.Application, opts ...igorm.Option) error {
	err := s.Session(opts...).WithContext(ctx).Create(in).Error
	if err != nil {
		return ormhelper.WrapErr(err)
	}
	return nil
}

func (s *Application) Get(ctx context.Context, id string, opts ...igorm.Option) (*models.Application, error) {
	var out models.Application
	err := s.Session(opts...).WithContext(ctx).Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, ormhelper.WrapErr(err)
	}
	return &out, nil
}

func (s *Application) CountByName(ctx context.Context, name string, opts ...igorm.Option) (int64, error) {
	var count int64
	err := s.Session(opts...).WithContext(ctx).Where("name = ?", name).Count(&count).Error
	if err != nil {
		return 0, ormhelper.WrapErr(err)
	}
	return count, nil
}

func (s *Application) Update(ctx context.Context, in *models.Application, opts ...igorm.Option) error {
	err := s.Session(opts...).WithContext(ctx).Updates(in).Error
	if err != nil {
		return ormhelper.WrapErr(err)
	}
	return nil
}

func (s *Application) Delete(ctx context.Context, id string, opts ...igorm.Option) error {
	err := s.Session(opts...).WithContext(ctx).Delete(&models.Application{
		ID: id,
	}).Error
	if err != nil {
		return ormhelper.WrapErr(err)
	}
	return nil
}

func (s *Application) Find(ctx context.Context, size int, offset string, opts ...igorm.Option) ([]*models.Application, string, error) {
	var (
		out        []*models.Application
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
