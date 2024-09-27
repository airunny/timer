package dao

import (
	"context"

	"github.com/airunny/timer/internal/models"

	"github.com/airunny/wiki-go-tools/igorm"
	"github.com/airunny/wiki-go-tools/ormhelper"
	"gorm.io/gorm"
)

type Token struct {
	db *gorm.DB
}

func NewToken(db *gorm.DB) *Token {
	return &Token{
		db: db,
	}
}

func (s *Token) Token(ctx context.Context, opts ...igorm.Option) *gorm.DB {
	return igorm.NewOptions(s.db, opts...).Session().WithContext(ctx)
}

func (s *Token) Add(ctx context.Context, in *models.Token, opts ...igorm.Option) error {
	err := s.Token(ctx, opts...).Create(in).Error
	if err != nil {
		return ormhelper.WrapErr(err)
	}
	return nil
}
