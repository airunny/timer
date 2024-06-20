package dao

import (
	"context"
	"math"

	"github.com/airunny/wiki-go-tools/igorm"
	"github.com/airunny/wiki-go-tools/ormhelper"
	redis "github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"timer/internal/models"
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

func (s *User) Create(ctx context.Context, in *models.User, opts ...igorm.Option) (int64, error) {
	err := s.Session(ctx, opts...).Create(in).Error
	if err != nil {
		return 0, ormhelper.WrapErr(err)
	}

	return in.ID, nil
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

func (s *User) Get(ctx context.Context, id int64, opts ...igorm.Option) (*models.User, error) {
	var out *models.User

	err := s.Session(ctx, opts...).
		Where("id = ?", id).
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

func (s *User) FindByOffset(ctx context.Context, offset int64, size int, opts ...igorm.Option) ([]*models.User, int64, error) {
	if offset <= 0 {
		offset = math.MaxInt64
	}

	var out []*models.User
	err := s.Session(ctx, opts...).
		Where("id < ?", offset). // 从offset处开始取
		Order("id desc").        // 如果要使用offset，则必须要按照ID进行排序
		Limit(size + 1).         // 这里多取一条
		Find(&out).Error
	if err != nil {
		return nil, 0, ormhelper.WrapErr(err)
	}

	var (
		count      = len(out)
		nextOffset int64
	)

	if count > 1 && count > size {
		nextOffset = out[count-2].ID
	}

	if count > size {
		out = out[:count-1]
	}

	return out, nextOffset, nil
}
