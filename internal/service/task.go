package service

import (
	"context"
	"time"

	v1 "github.com/airunny/timer/api/timer/v1"
	innErr "github.com/airunny/timer/errors"
	"github.com/airunny/timer/internal/models"
	"github.com/go-kratos/kratos/v2/log"
)

func (s *Service) ListTask(ctx context.Context, in *v1.ListTaskRequest) (*v1.ListTaskReply, error) {
	if in.Start >= in.End {
		return nil, innErr.WithMessage(innErr.ErrBadRequest, "invalid start time")
	}

	if in.Size <= 0 {
		in.Size = 10
	}

	var (
		l     = log.Context(ctx)
		start = time.Unix(in.Start, 0)
		end   = time.Unix(in.End, 0)
	)

	values, offset, err := s.task.FindByTime(ctx, in.Offset, start, end, int(in.Size), in.Status)
	if err != nil {
		l.Errorf("imerCallback.FindByTimerId Err:%v", err)
		return nil, err
	}

	items := make([]*v1.Task, 0, len(values))
	for _, value := range values {
		items = append(items, s.taskToGRPC(value))
	}

	return &v1.ListTaskReply{
		Items:  items,
		Offset: offset,
	}, nil
}

func (s *Service) ListTimerTask(ctx context.Context, in *v1.ListTimerTaskRequest) (*v1.ListTimerTaskReply, error) {
	if in.Size <= 0 {
		in.Size = 10
	}
	l := log.Context(ctx)

	values, offset, err := s.task.FindByTimerId(ctx, in.Id, in.Offset, int(in.Size), in.Status)
	if err != nil {
		l.Errorf("imerCallback.FindByTimerId Err:%v", err)
		return nil, err
	}

	items := make([]*v1.Task, 0, len(values))
	for _, value := range values {
		items = append(items, s.taskToGRPC(value))
	}

	return &v1.ListTimerTaskReply{
		Items:  items,
		Offset: offset,
	}, nil
}

func (s *Service) taskToGRPC(in *models.Task) *v1.Task {
	return &v1.Task{
		Id:           in.ID,
		TimerId:      in.TimerId,
		RunTime:      in.RunTime.Unix(),
		Request:      in.Request,
		Response:     in.Response,
		Status:       in.Status,
		Retry:        in.Retry,
		FailedReason: in.FailedReason,
		CreatedAt:    in.CreatedAt.Unix(),
		UpdatedAt:    in.UpdatedAt.Unix(),
	}
}
