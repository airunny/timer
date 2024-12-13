package dao

import (
	"context"

	v1 "github.com/airunny/timer/api/timer/v1"
	"github.com/airunny/timer/internal/models"

	"github.com/airunny/wiki-go-tools/igorm"
	"github.com/airunny/wiki-go-tools/ormhelper"
	redis "github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type User struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewUser(mysql *gorm.DB, redisClient *redis.Client) *User {
	return &User{
		db:    mysql,
		redis: redisClient,
	}
}

func (s *User) Session(ctx context.Context, opts ...igorm.Option) *gorm.DB {
	return igorm.NewOptions(s.db, opts...).Session().WithContext(ctx)
}

func (s *User) Add(ctx context.Context, in *models.User, opts ...igorm.Option) error {
	err := s.Session(ctx, opts...).Create(in).Error
	if err != nil {
		return ormhelper.WrapErr(err)
	}

	return nil
}

func (s *User) CountByName(ctx context.Context, name string, opts ...igorm.Option) (int64, error) {
	var count int64
	err := s.Session(ctx, opts...).
		Model(&models.User{}).
		Where("name = ?", name).
		Count(&count).Error
	if err != nil {
		return 0, ormhelper.WrapErr(err)
	}
	return count, nil
}

func (s *User) CountByNameWithInclude(ctx context.Context, name, excludeId string, opts ...igorm.Option) (int64, error) {
	var count int64
	err := s.Session(ctx, opts...).
		Model(&models.User{}).
		Where("name = ? and id != ?", name, excludeId).
		Count(&count).Error
	if err != nil {
		return 0, ormhelper.WrapErr(err)
	}
	return count, nil
}

func (s *User) Delete(ctx context.Context, id int64, opts ...igorm.Option) error {
	err := s.Session(ctx, opts...).Model(&models.User{}).Delete("id = ?", id).Error
	if err != nil {
		return ormhelper.WrapErr(err)
	}
	return nil
}

func (s *User) Update(ctx context.Context, in *models.User, opts ...igorm.Option) error {
	err := s.Session(ctx, opts...).Updates(in).Error
	if err != nil {
		return ormhelper.WrapErr(err)
	}
	return nil
}

func (s *User) UpdateStatus(ctx context.Context, id string, status v1.UserStatus, opts ...igorm.Option) error {
	err := s.Session(ctx, opts...).
		Model(&models.User{}).
		Where("id = ?", id).
		Update("status", status).Error
	if err != nil {
		return ormhelper.WrapErr(err)
	}
	return nil
}

func (s *User) Get(ctx context.Context, id string, opts ...igorm.Option) (*models.User, error) {
	var out *models.User
	err := s.Session(ctx, opts...).
		Where("id = ?", id).
		First(&out).Error
	if err != nil {
		return nil, ormhelper.WrapErr(err)
	}
	return out, nil
}

func (s *User) GetByName(ctx context.Context, name, password string, opts ...igorm.Option) (*models.User, error) {
	var out *models.User
	err := s.Session(ctx, opts...).
		Where("name = ? AND password = ?", name, password).
		First(&out).Error
	if err != nil {
		return nil, ormhelper.WrapErr(err)
	}
	return out, nil
}

func (s *User) FindByPage(ctx context.Context, page, size int, opts ...igorm.Option) ([]*models.User, error) {
	var out []*models.User
	err := s.Session(ctx, opts...).
		Order("id desc").
		Offset((page - 1) * size).
		Limit(size).
		Find(&out).Error
	if err != nil {
		return nil, ormhelper.WrapErr(err)
	}
	return out, nil
}

func (s *User) FindByOffset(ctx context.Context, size int, offset string, opts ...igorm.Option) ([]*models.User, string, error) {
	var (
		out        []*models.User
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
