package dao

import (
	"errors"
	"time"

	"github.com/airunny/timer/api/common"
	"github.com/airunny/timer/internal/models"
	"github.com/airunny/wiki-go-tools/env"
	"github.com/airunny/wiki-go-tools/igorm"
	"github.com/airunny/wiki-go-tools/iredis"
	"github.com/go-kratos/kratos/v2/log"
	redis "github.com/go-redis/redis/v8"
	goCache "github.com/liyanbing/go-cache"
	redisCache "github.com/liyanbing/go-cache/cacher/redis"
	"gorm.io/gorm"
)

func NewMySQL(c *common.DataConfig, logger log.Logger) (*gorm.DB, func(), error) {
	database := c.Database
	if database == nil {
		return nil, nil, errors.New("empty database config")
	}

	db, closer, err := igorm.NewGORM(&igorm.Config{
		LogLevel:    int(database.Level),
		MaxOpen:     int(database.MaxOpen),
		MaxIdle:     int(database.MaxIdle),
		MaxLifeTime: time.Duration(database.MaxLifeTimeSeconds) * time.Second,
		Source:      c.Database.Source,
	}, logger)
	if err != nil {
		return nil, nil, err
	}

	err = db.AutoMigrate(
		&models.Application{},
		&models.Timer{},
		&models.TimerCallback{},
		&models.User{},
		&models.Token{},
	)
	if err != nil {
		return nil, nil, err
	}

	return db, func() {
		_ = closer.Close()
	}, nil
}

func NewRedis(c *common.DataConfig, logger log.Logger) (*redis.Client, func(), error) {
	db, err := iredis.NewClient(&iredis.Config{
		Address:  c.Redis.Address,
		Password: c.Redis.Password,
		DB:       int(c.Redis.Db),
		MaxIdle:  int(c.Redis.MaxIdle),
	}, logger)
	if err != nil {
		return nil, nil, err
	}

	return db, func() {
		_ = db.Close()
	}, nil
}

func NewRedisCache(redisClient *redis.Client) (goCache.Cache, error) {
	cache := redisCache.NewRedisCache(redisClient)
	cache.SetNamespace(env.GetServiceName())
	return cache, nil
}
