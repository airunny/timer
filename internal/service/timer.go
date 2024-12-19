package service

import (
	"context"
	"time"

	v1 "github.com/airunny/timer/api/timer/v1"
	"github.com/airunny/timer/errors"
	"github.com/airunny/timer/internal/models"
	"github.com/airunny/wiki-go-tools/objectid"
	"github.com/go-kratos/kratos/v2/log"
)

func (s *Service) AddTimer(ctx context.Context, in *v1.AddTimerRequest) (*v1.AddTimerReply, error) {
	var (
		l   = log.Context(ctx)
		now = time.Now()
	)

	// 不能小于1分钟内
	if in.ExpireAt <= now.Add(time.Minute).Unix() {
		return nil, errors.WithMessage(errors.ErrBadRequest, "the interval cannot be less than one minute")
	}

	if in.CallbackType <= 0 {
		return nil, errors.ErrBadRequest
	}

	if in.CallbackType == v1.CallbackType_KAFKA && s.producer == nil {
		return nil, errors.WithMessage(errors.ErrBadRequest, "kafka not supported ,please add kafka configuration")
	}

	// TODO 参数校验
	newId := objectid.ObjectID()
	err := s.timer.Add(ctx, &models.Timer{
		ID:              newId,
		AppId:           in.AppId,
		Name:            in.Name,
		Type:            in.Type,
		CallbackType:    in.CallbackType,
		CallbackAddress: in.CallbackAddress,
		CronExpr:        in.CronExpr,
		ExpireAt:        in.ExpireAt,
		Attach:          in.Attach,
		Status:          in.Status,
		Extra: &models.TimerExtra{
			Params: in.Params,
		},
	})
	if err != nil {
		l.Errorf("timer.Add Err:%v", err)
		return nil, err
	}

	// TODO to redis
	return &v1.AddTimerReply{
		Id: newId,
	}, nil
}

func (s *Service) GetTimer(ctx context.Context, in *v1.GetTimerRequest) (*v1.Timer, error) {
	l := log.Context(ctx)

	timer, err := s.timer.Get(ctx, in.Id)
	if err != nil {
		l.Errorf("timer.Get Err:%v", err)
		return nil, err
	}

	ttl := time.Now().Unix() - timer.ExpireAt
	if ttl < 0 {
		ttl = -1
	}
	return s.timerToGRPC(timer), nil
}

func (s *Service) UpdateTimerStatus(ctx context.Context, in *v1.UpdateTimerStatusRequest) (*v1.UpdateTimerStatusReply, error) {
	l := log.Context(ctx)
	err := s.timer.UpdateStatus(ctx, in.Id, in.Status)
	if err != nil {
		l.Errorf("timer.UpdateStatus Err:%v", err)
		return nil, err
	}
	return &v1.UpdateTimerStatusReply{}, nil
}

func (s *Service) DeleteTimer(ctx context.Context, in *v1.DeleteTimerRequest) (*v1.DeleteTimerReply, error) {
	l := log.Context(ctx)
	err := s.timer.Delete(ctx, in.Id)
	if err != nil {
		l.Errorf("timer.UpdateStatus Err:%v", err)
		return nil, err
	}
	return &v1.DeleteTimerReply{}, nil
}

func (s *Service) ReplayTimer(ctx context.Context, in *v1.ReplayTimerRequest) (*v1.ReplayTimerReply, error) {
	// TODO
	return nil, nil
}

func (s *Service) ListTimer(ctx context.Context, in *v1.ListTimerRequest) (*v1.ListTimerReply, error) {
	l := log.Context(ctx)

	if in.Size <= 0 {
		in.Size = 10
	}

	values, newOffset, err := s.timer.Find(ctx, int(in.Size), in.Offset)
	if err != nil {
		l.Errorf("timer.Find Err:%v", err)
		return nil, err
	}

	items := make([]*v1.Timer, 0, len(values))
	for _, value := range values {
		items = append(items, s.timerToGRPC(value))
	}

	return &v1.ListTimerReply{
		Items:  items,
		Offset: newOffset,
	}, nil
}

func (s *Service) ListTimerCallback(ctx context.Context, in *v1.ListTimerCallbackRequest) (*v1.ListTimerCallbackReply, error) {
	l := log.Context(ctx)
	values, offset, err := s.timerCallback.FindByTimerId(ctx, in.Id, in.Offset, int(in.Size))
	if err != nil {
		l.Errorf("imerCallback.FindByTimerId Err:%v", err)
		return nil, err
	}

	items := make([]*v1.TimerCallback, 0, len(values))
	for _, value := range values {
		items = append(items, &v1.TimerCallback{
			Id:             value.ID,
			RequestContent: value.RequestContent,
			ReplyContent:   value.ReplyContent,
			FailedReason:   value.FailedReason,
			RetryCount:     value.RetryCount,
			Status:         value.Status,
			CreatedAt:      value.CreatedAt.Unix(),
		})
	}

	return &v1.ListTimerCallbackReply{
		Items:  items,
		Offset: offset,
	}, nil
}

func (s *Service) timerToGRPC(in *models.Timer) *v1.Timer {
	ttl := time.Now().Unix() - in.ExpireAt
	if ttl < 0 {
		ttl = -1
	}

	return &v1.Timer{
		Id:                    in.ID,
		AppId:                 in.AppId,
		Name:                  in.Name,
		Type:                  in.Type,
		CallbackType:          in.CallbackType,
		CallbackAddress:       in.CallbackAddress,
		CronExpr:              in.CronExpr,
		ExpireAt:              in.ExpireAt,
		Attach:                in.Attach,
		Ttl:                   ttl,
		Fail:                  in.Fail,
		Success:               in.Success,
		CallbackFailedReasons: in.Extra.FailReason,
		Params:                in.Extra.Params,
		Status:                in.Status,
	}
}
