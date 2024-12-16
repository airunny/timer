package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
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
		newId  = objectid.ObjectID()
		now    = time.Now()
		secret = s.generateApplicationSecret(newId)
		newApp = &models.Application{
			ID:            newId,
			Name:          in.Name,
			Description:   in.Description,
			Secret:        secret,
			Authorization: in.Authentication,
			Status:        in.Status,
			CreatedAt:     now,
			UpdatedAt:     now,
		}
	)

	err = s.application.Add(ctx, newApp)
	if err != nil {
		l.Errorf("application.Add Err:%v", err)
		return nil, err
	}
	return s.applicationToGRPC(newApp), nil
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
	var (
		l      = log.Context(ctx)
		secret = s.generateApplicationSecret(in.Id)
	)
	err := s.application.UpdateSecret(ctx, in.Id, secret)
	if err != nil {
		l.Errorf("UpdateSecret Err:%v", err)
		return nil, err
	}
	return &v1.GenApplicationSecretReply{
		Secret: secret,
	}, nil
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
		return nil, innerErr.WithMessage(innerErr.ErrBadRequest, "name can not be empty")
	}

	nameCount, err := s.application.CountByNameWithInclude(ctx, in.Name, in.Id)
	if err != nil {
		l.Errorf("CountByName Err:%v", err)
		return nil, err
	}

	if nameCount > 0 {
		return nil, innerErr.WithMessage(innerErr.ErrBadRequest, "name already exists")
	}

	err = s.application.Update(ctx, &models.Application{
		ID:            in.Id,
		Name:          in.Name,
		Description:   in.Description,
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

func (s *Service) generateApplicationSecret(id string) string {
	h := md5.New()
	h.Write([]byte(fmt.Sprintf("%s:%d:%d", id, time.Now().Unix(), rand.Int63())))
	return hex.EncodeToString(h.Sum(nil))
}
