package dao

import (
	"context"

	"github.com/airunny/timer/internal/models"
	"github.com/airunny/wiki-go-tools/igorm"
	"github.com/airunny/wiki-go-tools/ormhelper"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Token struct {
	db *gorm.DB
}

func NewToken(db *gorm.DB) *Token {
	return &Token{
		db: db,
	}
}

func (s *Token) session(ctx context.Context, opts ...igorm.Option) *gorm.DB {
	return igorm.NewOptions(s.db, opts...).Session().WithContext(ctx)
}

func (s *Token) Add(ctx context.Context, in *models.Token, opts ...igorm.Option) error {
	err := s.session(ctx, opts...).Create(in).Error
	if err != nil {
		return ormhelper.WrapErr(err)
	}
	return nil
}

func (s *Token) GetByRefreshToken(ctx context.Context, in string, opts ...igorm.Option) (*models.Token, error) {
	var out models.Token
	err := s.session(ctx, opts...).Where("refresh_token = ?", in).First(&out).Error
	if err != nil {
		return nil, ormhelper.WrapErr(err)
	}
	return &out, nil
}

func (s *Token) Upsert(ctx context.Context, in *models.Token, opts ...igorm.Option) error {
	err := s.session(ctx, opts...).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "user_id"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"access_token":       in.AccessToken,
				"expires_in":         in.ExpiresIn,
				"refresh_token":      in.RefreshToken,
				"refresh_expires_in": in.RefreshExpiresIn,
			}),
		}).
		Create(in).Error
	if err != nil {
		return ormhelper.WrapErr(err)
	}
	return nil
}
