package service

import (
	"context"
	"strings"

	"github.com/airunny/wiki-go-tools/objectid"
	"github.com/go-kratos/kratos/v2/log"
	v1 "timer/api/timer/v1"
	innerErr "timer/errors"
	"timer/internal/models"
)

func (s *Service) AddApplication(ctx context.Context, in *v1.AddApplicationRequest) (*v1.Application, error) {
	l := log.Context(ctx)

	if strings.TrimSpace(in.Name) == "" {
		return nil, innerErr.WithMessage(innerErr.ErrBadRequest, "name 不能为空")
	}

	nameCount, err := s.application.CountByName(ctx, in.Name)
	if err != nil {
		l.Errorf("CountByName Err:%v", err)
		return nil, err
	}

	if nameCount > 0 {
		return nil, innerErr.WithMessage(innerErr.ErrBadRequest, "name 已存在")
	}

	newID := objectid.ObjectID()
	err = s.application.Add(ctx, &models.Application{
		ID:            newID,
		Name:          in.Name,
		Description:   in.Description,
		Authorization: in.Authentication,
	})
	if err != nil {
		l.Errorf("application.Add Err:%v", err)
		return nil, err
	}

	return &v1.Application{
		Id: newID,
	}, nil
}

func (s *Service) GetApplication(ctx context.Context, in *v1.GetApplicationRequest) (*v1.Application, error) {
	l := log.Context(ctx)

	application, err := s.application.Get(ctx, in.Id)
	if err != nil {
		l.Errorf("application.Get Err:%v", err)
		return nil, err
	}

	return &v1.Application{
		Id:             application.ID,
		Name:           application.Name,
		Description:    application.Description,
		CreatedAt:      application.CreatedAt.Unix(),
		Secret:         application.Secret,
		UpdatedAt:      application.UpdatedAt.Unix(),
		Authentication: application.Authorization,
	}, nil
}

func (s *Service) DeleteApplication(ctx context.Context, in *v1.DeleteApplicationRequest) (*v1.DeleteApplicationReply, error) {

}

func (s *Service) UpdateApplication(ctx context.Context, in *v1.UpdateApplicationRequest) (*v1.UpdateApplicationReply, error) {

}

func (s *Service) ListApplication(ctx context.Context, in *v1.ListApplicationRequest) (*v1.ListApplicationReply, error) {

}
