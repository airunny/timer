package service

import (
	"context"
	"strings"

	v1 "github.com/airunny/timer/api/timer/v1"
	innerErr "github.com/airunny/timer/errors"
	"github.com/airunny/timer/internal/models"
	"github.com/airunny/wiki-go-tools/objectid"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) AddUser(ctx context.Context, in *v1.AddUserRequest) (*v1.AddUserReply, error) {
	l := log.Context(ctx)

	if strings.TrimSpace(in.Name) == "" {
		return nil, innerErr.WithMessage(innerErr.ErrBadRequest, "name can not be empty")
	}

	nameCount, err := s.user.CountByName(ctx, in.Name)
	if err != nil {
		l.Errorf("CountByName Err:%v", err)
		return nil, err
	}

	if nameCount > 0 {
		return nil, innerErr.WithMessage(innerErr.ErrBadRequest, "name already exists")
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		l.Errorf("GenerateFromPassword Err:%v", err)
		return nil, err
	}

	newId := objectid.ObjectID()
	err = s.user.Add(ctx, &models.User{
		ID:       newId,
		Name:     in.Name,
		Password: string(pwd),
		Status:   in.Status,
		Role:     in.Role,
	})
	if err != nil {
		l.Errorf("user.Add Err:%v", err)
		return nil, err
	}
	return &v1.AddUserReply{
		Id: newId,
	}, nil
}

func (s *Service) GetUser(ctx context.Context, in *v1.GetUserRequest) (*v1.User, error) {
	l := log.Context(ctx)

	user, err := s.user.Get(ctx, in.Id)
	if err != nil {
		l.Errorf("user.Get Err:%v", err)
		return nil, err
	}

	return &v1.User{
		Id:        user.ID,
		Name:      user.Name,
		Status:    user.Status,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
	}, nil
}

func (s *Service) UpdateUserStatus(ctx context.Context, in *v1.UpdateUserStatusRequest) (*v1.UpdateUserStatusReply, error) {
	l := log.Context(ctx)
	err := s.user.UpdateStatus(ctx, in.Id, in.Status)
	if err != nil {
		l.Errorf("UpdateStatus Err:%v", err)
		return nil, err
	}

	return &v1.UpdateUserStatusReply{
		Status: in.Status,
	}, nil
}

func (s *Service) UpdateUserPassword(ctx context.Context, in *v1.UpdateUserPasswordRequest) (*v1.UpdateUserPasswordReply, error) {
	if in.OldPassword == in.NewPassword {
		return nil, innerErr.WithMessage(innerErr.ErrBadRequest, "the new password cannot be the same as the old password")
	}
	l := log.Context(ctx)

	user, err := s.user.Get(ctx, in.Id)
	if err != nil {
		l.Errorf("user.Get Err:%v", err)
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.OldPassword))
	if err != nil {
		return nil, innerErr.WithMessage(innerErr.ErrBadRequest, "old password is incorrect")
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(in.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		l.Errorf("GenerateFromPassword Err:%v", err)
		return nil, err
	}

	err = s.user.UpdatePassword(ctx, in.Id, string(pwd))
	if err != nil {
		l.Errorf("UpdatePassword Err:%v", err)
		return nil, err
	}
	return &v1.UpdateUserPasswordReply{}, nil
}

func (s *Service) ListUser(ctx context.Context, in *v1.ListUserRequest) (*v1.ListUserReply, error) {
	l := log.Context(ctx)

	if in.Size <= 0 {
		in.Size = 10
	}

	values, newOffset, err := s.user.FindByOffset(ctx, int(in.Size), in.Offset)
	if err != nil {
		l.Errorf("Find Err:%v", err)
		return nil, err
	}

	items := make([]*v1.User, 0, len(values))
	for _, value := range values {
		items = append(items, s.userToGRPC(value))
	}

	return &v1.ListUserReply{
		Items:  items,
		Offset: newOffset,
	}, nil
}

func (s *Service) DeleteUser(ctx context.Context, in *v1.DeleteUserRequest) (*v1.DeleteUserReply, error) {
	l := log.Context(ctx)
	if in.Id == adminPassword {
		return nil, innerErr.WithMessage(innerErr.ErrBadRequest, "admin cannot be deleted")
	}

	err := s.user.Delete(ctx, in.Id)
	if err != nil {
		l.Errorf("user.Delete Err:%v", err)
		return nil, err
	}
	return &v1.DeleteUserReply{}, nil
}

func (s *Service) userToGRPC(in *models.User) *v1.User {
	return &v1.User{
		Id:        in.ID,
		Name:      in.Name,
		Role:      in.Role,
		CreatedAt: in.CreatedAt.Unix(),
		UpdatedAt: in.UpdatedAt.Unix(),
		Status:    in.Status,
	}
}
