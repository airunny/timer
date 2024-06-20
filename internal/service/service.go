package service

import (
	"context"

	"github.com/google/wire"
	"timer/api/common"
	v1 "timer/api/timer/v1"
	"timer/internal/conf"
	"timer/internal/dao"
)

var ProviderSet = wire.NewSet(NewService)

type Service struct {
	v1.UnimplementedServiceServer
	user        *dao.User
	application *dao.Application
	timerRecord *dao.TimerRecord
}

func NewService(
	business *conf.Business,
	user *dao.User,
	application *dao.Application,
	timerRecord *dao.TimerRecord,
) *Service {
	return &Service{
		UnimplementedServiceServer: v1.UnimplementedServiceServer{},
		user:                       user,
		application:                application,
		timerRecord:                timerRecord,
	}
}

func (s *Service) Healthy(_ context.Context, _ *common.EmptyRequest) (*common.HealthyReply, error) {
	return &common.HealthyReply{
		Status: common.HealthyStatus_HealthyStatusSERVING,
	}, nil
}
