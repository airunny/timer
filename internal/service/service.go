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
	"github.com/airunny/wiki-go-tools/ormhelper"
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
	business      *conf.Business
	user          *dao.User
	application   *dao.Application
	timer         *dao.TimerRecord
	timerCallback *dao.TimerCallback
	token         *dao.Token
	producer      *ikafka.Producer
	JWT           *jwt.JWT
}

func NewService(
	business *conf.Business,
	user *dao.User,
	application *dao.Application,
	timer *dao.TimerRecord,
	timerCallback *dao.TimerCallback,
	token *dao.Token,
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

	jwtCli, err := jwt.NewJWT(business.JWT)
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

	s := &Service{
		UnimplementedServiceServer: v1.UnimplementedServiceServer{},
		business:                   business,
		user:                       user,
		application:                application,
		timer:                      timer,
		timerCallback:              timerCallback,
		token:                      token,
		producer:                   producer,
		JWT:                        jwtCli,
	}
	return s, func() {

	}
}

func (s *Service) Healthy(_ context.Context, _ *common.EmptyRequest) (*common.HealthyReply, error) {
	return &common.HealthyReply{
		Status: common.HealthyStatus_HealthyStatusSERVING,
	}, nil
}
