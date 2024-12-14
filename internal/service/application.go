package service

import (
	"context"
	"strings"
	"time"

	v1 "github.com/airunny/timer/api/timer/v1"
	innerErr "github.com/airunny/timer/errors"
	"github.com/airunny/timer/internal/models"
	"github.com/airunny/wiki-go-tools/objectid"
	"github.com/go-kratos/kratos/v2/log"
)

func (s *Service) AddApplication(ctx context.Context, in *v1.AddApplicationRequest) (*v1.Application, error) {
	l := log.Context(ctx)

	if strings.TrimSpace(in.Name) == "" {
		return nil, innerErr.WithMessage(innerErr.ErrBadRequest, "name can not be empty")
	}

	nameCount, err := s.application.CountByName(ctx, in.Name)
	if err != nil {
		l.Errorf("CountByName Err:%v", err)
		return nil, err
	}

	if nameCount > 0 {
		return nil, innerErr.WithMessage(innerErr.ErrBadRequest, "name already exists")
	}

	var (
		newId = objectid.ObjectID()
		now   = time.Now()
	)

	err = s.application.Add(ctx, &models.Application{
		ID:            newId,
		Name:          in.Name,
		Description:   in.Description,
		Secret:        "", // TODO
		Authorization: in.Authentication,
		Status:        in.Status,
	})
	if err != nil {
		l.Errorf("application.Add Err:%v", err)
		return nil, err
	}

	return &v1.Application{
		Id:             newId,
		Name:           in.Name,
		Description:    in.Description,
		CreatedAt:      now.Unix(),
		Secret:         "", // TODO
		UpdatedAt:      now.Unix(),
		Authentication: in.Authentication,
	}, nil
}

func (s *Service) GetApplication(ctx context.Context, in *v1.GetApplicationRequest) (*v1.Application, error) {
	l := log.Context(ctx)

	application, err := s.application.Get(ctx, in.Id)
	if err != nil {
		l.Errorf("application.Get Err:%v", err)
		return nil, err
	}

	return s.applicationToGRPC(application), nil
}

func (s *Service) UpdateApplicationStatus(ctx context.Context, in *v1.UpdateApplicationStatusRequest) (*v1.UpdateApplicationStatusReply, error) {
	l := log.Context(ctx)
	err := s.application.UpdateStatus(ctx, in.Id, in.Status)
	if err != nil {
		l.Errorf("UpdateStatus Err:%v", err)
		return nil, err
	}

	return &v1.UpdateApplicationStatusReply{
		Status: in.Status,
	}, nil
}

func (s *Service) GenApplicationSecret(ctx context.Context, in *v1.GenApplicationSecretRequest) (*v1.GenApplicationSecretReply, error) {
	l := log.Context(ctx)
	// TODO
	err := s.application.UpdateSecret(ctx, in.Id, "")
	if err != nil {
		l.Errorf("UpdateSecret Err:%v", err)
		return nil, err
	}

	return &v1.GenApplicationSecretReply{}, nil
}

func (s *Service) DeleteApplication(ctx context.Context, in *v1.DeleteApplicationRequest) (*v1.DeleteApplicationReply, error) {
	l := log.Context(ctx)
	err := s.application.Delete(ctx, in.Id)
	if err != nil {
		l.Errorf("Delete Err:%v", err)
		return nil, err
	}

	return &v1.DeleteApplicationReply{}, nil
}

func (s *Service) UpdateApplication(ctx context.Context, in *v1.UpdateApplicationRequest) (*v1.UpdateApplicationReply, error) {
	l := log.Context(ctx)

	if strings.TrimSpace(in.Name) == "" {
		return nil, innerErr.WithMessage(innerErr.ErrBadRequest, "name 不能为空")
	}

	nameCount, err := s.application.CountByNameWithInclude(ctx, in.Name, in.Id)
	if err != nil {
		l.Errorf("CountByName Err:%v", err)
		return nil, err
	}

	if nameCount > 0 {
		return nil, innerErr.WithMessage(innerErr.ErrBadRequest, "name 已存在")
	}

	err = s.application.Update(ctx, &models.Application{
		ID:            in.Id,
		Name:          in.Name,
		Description:   in.Description,
		Secret:        "", // TODO
		Authorization: in.Authentication,
		Status:        in.Status,
	})
	if err != nil {
		l.Errorf("Update Err:%v", err)
		return nil, err
	}

	return &v1.UpdateApplicationReply{}, nil
}

func (s *Service) ListApplication(ctx context.Context, in *v1.ListApplicationRequest) (*v1.ListApplicationReply, error) {
	l := log.Context(ctx)

	if in.Size <= 0 {
		in.Size = 10
	}

	values, newOffset, err := s.application.Find(ctx, int(in.Size), in.Offset)
	if err != nil {
		l.Errorf("Find Err:%v", err)
		return nil, err
	}

	items := make([]*v1.Application, 0, len(values))
	for _, value := range values {
		items = append(items, s.applicationToGRPC(value))
	}

	return &v1.ListApplicationReply{
		Items:  items,
		Offset: newOffset,
	}, nil
}

func (s *Service) applicationToGRPC(in *models.Application) *v1.Application {
	return &v1.Application{
		Id:             in.ID,
		Name:           in.Name,
		Description:    in.Description,
		CreatedAt:      in.CreatedAt.Unix(),
		Secret:         in.Secret,
		UpdatedAt:      in.UpdatedAt.Unix(),
		Authentication: in.Authorization,
		Status:         in.Status,
	}
}
