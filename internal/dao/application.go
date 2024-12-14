package dao

import (
	"context"

	v1 "github.com/airunny/timer/api/timer/v1"
	"github.com/airunny/timer/internal/models"
	"github.com/airunny/wiki-go-tools/igorm"
	"github.com/airunny/wiki-go-tools/ormhelper"
	"gorm.io/gorm"
)

type Application struct {
	db *gorm.DB
}

func NewApplication(db *gorm.DB) *Application {
	return &Application{
		db: db,
	}
}

func (s *Application) Session(ctx context.Context, opts ...igorm.Option) *gorm.DB {
	return igorm.NewOptions(s.db, opts...).Session().WithContext(ctx)
}

func (s *Application) Add(ctx context.Context, in *models.Application, opts ...igorm.Option) error {
	err := s.Session(ctx, opts...).
		Create(in).Error
	if err != nil {
		return ormhelper.WrapErr(err)
	}
	return nil
}

func (s *Application) Get(ctx context.Context, id string, opts ...igorm.Option) (*models.Application, error) {
	var out models.Application
	err := s.Session(ctx, opts...).
		Where("id = ?", id).
		First(&out).Error
	if err != nil {
		return nil, ormhelper.WrapErr(err)
	}
	return &out, nil
}

func (s *Application) CountByName(ctx context.Context, name string, opts ...igorm.Option) (int64, error) {
	var count int64
	err := s.Session(ctx, opts...).
		Model(&models.Application{}).
		Where("name = ?", name).
		Count(&count).Error
	if err != nil {
		return 0, ormhelper.WrapErr(err)
	}
	return count, nil
}

func (s *Application) CountByNameWithInclude(ctx context.Context, name, excludeId string, opts ...igorm.Option) (int64, error) {
	var count int64
	err := s.Session(ctx, opts...).
		Model(&models.Application{}).
		Where("name = ? and id != ?", name, excludeId).
		Count(&count).Error
	if err != nil {
		return 0, ormhelper.WrapErr(err)
	}
	return count, nil
}

func (s *Application) Update(ctx context.Context, in *models.Application, opts ...igorm.Option) error {
	err := s.Session(ctx, opts...).
		Select("name", "description", "authorization", "status").
		Updates(in).Error
	if err != nil {
		return ormhelper.WrapErr(err)
	}
	return nil
}

func (s *Application) UpdateStatus(ctx context.Context, id string, status v1.ApplicationStatus, opts ...igorm.Option) error {
	err := s.Session(ctx, opts...).
		Model(&models.Application{}).
		Where("id = ?", id).
		Update("status", status).Error
	if err != nil {
		return ormhelper.WrapErr(err)
	}
	return nil
}

func (s *Application) UpdateSecret(ctx context.Context, id, secret string, opts ...igorm.Option) error {
	err := s.Session(ctx, opts...).
		Model(&models.Application{}).
		Where("id = ?", id).
		Update("secret", secret).Error
	if err != nil {
		return ormhelper.WrapErr(err)
	}
	return nil
}

func (s *Application) Delete(ctx context.Context, id string, opts ...igorm.Option) error {
	err := s.Session(ctx, opts...).
		Delete(&models.Application{
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
		session    = s.Session(ctx, opts...).Order("id desc")
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
