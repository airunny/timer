package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/airunny/timer/api/common"
	v1 "github.com/airunny/timer/api/timer/v1"
	"github.com/airunny/timer/internal/conf"
	"github.com/airunny/timer/internal/dao"
	"github.com/airunny/timer/internal/models"
	"github.com/airunny/timer/pkg/ikafka"
	"github.com/airunny/timer/pkg/jwt"
	"github.com/airunny/timer/pkg/queue"
	"github.com/airunny/wiki-go-tools/locker"
	"github.com/airunny/wiki-go-tools/ormhelper"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewService)

const (
	adminUserId   = "000000000000000000000001"
	adminUserName = "timer"
	adminPassword = "$2a$10$gYhHCgjz9l307bYbfr0sFelCdl4jHX0cZhs8/hynfYlsNf8yWOiAu"
)

type Service struct {
	v1.UnimplementedServiceServer
	redisClient *redis.Client
	business    *conf.Business
	user        *dao.User
	application *dao.Application
	timer       *dao.TimerRecord
	task        *dao.Task
	token       *dao.Token
	producer    *ikafka.Producer
	JWT         *jwt.JWT
	locker      locker.Locker
	global      *dao.Global
	publisher   queue.Publisher
	consumer    queue.Consumer
}

func NewService(
	redisClient *redis.Client,
	business *conf.Business,
	user *dao.User,
	application *dao.Application,
	timer *dao.TimerRecord,
	task *dao.Task,
	token *dao.Token,
	global *dao.Global,
) (*Service, func()) {
	var (
		producer *ikafka.Producer
		err      error
	)

	if business.KAFKA != nil {
		producer, err = ikafka.NewProducer(business.KAFKA)
		if err != nil {
			panic(err)
		}
	}

	if business.JWT == nil {
		business.JWT = &jwt.Config{
			Key:    "timer:9d0xJlx!x914",
			Issuer: "timer",
		}
	}

	jwtClient, err := jwt.NewJWT(business.JWT)
	if err != nil {
		panic(err)
	}

	// 初始化管理员
	err = user.Add(context.Background(), &models.User{
		ID:       adminUserId,
		Name:     adminUserName,
		Password: adminPassword,
		Status:   v1.UserStatus_UserStatusOn,
		Role:     v1.UserRole_Admin,
	})
	if err != nil && !errors.Is(err, ormhelper.ErrDuplicateKey) {
		panic(fmt.Errorf("init admin user err %v", err))
	}

	l, err := locker.NewLockerWithRedis(redisClient)
	if err != nil {
		panic(err)
	}

	consumer := queue.NewConsumerWithRedis(redisClient)
	s := &Service{
		redisClient: redisClient,
		business:    business,
		user:        user,
		application: application,
		timer:       timer,
		task:        task,
		token:       token,
		producer:    producer,
		JWT:         jwtClient,
		locker:      l,
		publisher:   queue.NewPublishWithRedis(redisClient),
		consumer:    consumer,
		global:      global,
	}

	return s, func() {
		consumer.Close()
	}
}

func (s *Service) Healthy(_ context.Context, _ *common.EmptyRequest) (*common.HealthyReply, error) {
	return &common.HealthyReply{
		Status: common.HealthyStatus_HealthyStatusSERVING,
	}, nil
}
