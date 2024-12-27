package dao

import (
	"context"

	"github.com/airunny/timer/internal/models"
	"github.com/airunny/wiki-go-tools/igorm"
	"github.com/airunny/wiki-go-tools/ormhelper"
	"gorm.io/gorm"
)

type Global struct {
	db *gorm.DB
}

func NewGlobal(db *gorm.DB) *Global {
	return &Global{
		db: db,
	}
}

func (s *Global) session(ctx context.Context, opts ...igorm.Option) *gorm.DB {
	return igorm.NewOptions(s.db, opts...).Session().WithContext(ctx)
}

func (s *Global) Add(ctx context.Context, in *models.Global, opts ...igorm.Option) (int64, error) {
	err := s.session(ctx, opts...).Create(in).Error
	if err != nil {
		return 0, ormhelper.WrapErr(err)
	}
	return in.ID, nil
}

func (s *Global) GetByName(ctx context.Context, namespace, name string, opts ...igorm.Option) (*models.Global, error) {
	var out models.Global
	err := s.session(ctx, opts...).
		Where("namespace = ? and name = ?", namespace, name).
		First(&out).Error
	if err != nil {
		return nil, ormhelper.WrapErr(err)
	}
	return &out, nil
}

func (s *Global) UpsertByKey(ctx context.Context, in *models.Global, opts ...igorm.Option) error {
	err := s.session(ctx, opts...).
		Model(&models.Global{}).
		Where("namespace = ? and name = ?", in.Namespace, in.Name).
		Update("value", in.Value).Error
	if err != nil {
		return ormhelper.WrapErr(err)
	}
	return nil
}

func (s *Global) DeleteByKey(ctx context.Context, namespace, name string, opts ...igorm.Option) error {
	err := s.session(ctx, opts...).
		Delete(&models.Global{}, "namespace = ? and name = ?", namespace, name).Error
	if err != nil {
		return ormhelper.WrapErr(err)
	}
	return nil
}
