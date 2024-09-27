package service

import (
	"context"

	"github.com/airunny/timer/api/common"
	v1 "github.com/airunny/timer/api/timer/v1"
	"github.com/airunny/timer/internal/conf"
	"github.com/airunny/timer/internal/dao"
	"github.com/airunny/timer/pkg/ikafka"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewService)

type Service struct {
	v1.UnimplementedServiceServer
	user          *dao.User
	application   *dao.Application
	timer         *dao.TimerRecord
	timerCallback *dao.TimerCallback
	producer      *ikafka.Producer
}

func NewService(
	business *conf.Business,
	user *dao.User,
	application *dao.Application,
	timer *dao.TimerRecord,
	timerCallback *dao.TimerCallback,
) (*Service, func()) {
	var (
		producer *ikafka.Producer
		//err      error
	)

	//if business.KAFKA != nil {
	//	producer, err = ikafka.NewProducer(business.KAFKA)
	//	if err != nil {
	//		panic(err)
	//	}
	//}

	s := &Service{
		UnimplementedServiceServer: v1.UnimplementedServiceServer{},
		user:                       user,
		application:                application,
		timer:                      timer,
		timerCallback:              timerCallback,
		producer:                   producer,
	}

	return s, s.Close
}

func (s *Service) Healthy(_ context.Context, _ *common.EmptyRequest) (*common.HealthyReply, error) {
	return &common.HealthyReply{
		Status: common.HealthyStatus_HealthyStatusSERVING,
	}, nil
}

func (s *Service) Close() {
}
