package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	v1 "github.com/airunny/timer/api/timer/v1"
	innerErr "github.com/airunny/timer/errors"
	"github.com/airunny/timer/internal/models"
	"github.com/airunny/timer/pkg/jwt"
	"github.com/airunny/wiki-go-tools/ormhelper"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	l := log.Context(ctx)

	user, err := s.user.GetByName(ctx, in.Username)
	if err != nil && errors.Is(err, ormhelper.ErrNotFound) {
		l.Errorf("GetByName Err:%v", err)
		return nil, err
	}

	if err != nil {
		return nil, innerErr.WithMessage(innerErr.ErrBadRequest, "name or password incorrect")
	}

	if user.Status != v1.UserStatus_UserStatusOn {
		return nil, innerErr.WithMessage(innerErr.ErrBadRequest, "user not available")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password))
	if err != nil {
		return nil, innerErr.WithMessage(innerErr.ErrBadRequest, "name or password incorrect")
	}

	token, err := s.generateToken(ctx, user)
	if err != nil {
		l.Error("generateToken Err:%v", err)
		return nil, err
	}

	return &v1.LoginReply{
		AccessToken:        token.AccessToken,
		ExpiresIn:          token.ExpiresIn,
		RefreshToken:       token.RefreshToken,
		RefreshExpiresIn:   token.RefreshExpiresIn,
		NeedChangePassword: !user.ChangePassword,
	}, nil
}

func (s *Service) RefreshToken(ctx context.Context, in *v1.RefreshTokenRequest) (*v1.LoginReply, error) {
	if in.RefreshToken == "" {
		return nil, innerErr.ErrBadRequest
	}
	l := log.Context(ctx)

	token, err := s.token.GetByRefreshToken(ctx, in.RefreshToken)
	if err != nil && errors.Is(err, ormhelper.ErrNotFound) {
		l.Errorf("GetByRefreshToken Err:%v", err)
		return nil, err
	}

	if err != nil {
		return nil, innerErr.ErrResourceNotFound
	}

	if token.RefreshExpiresIn <= time.Now().Unix() {
		return nil, innerErr.ErrRefreshTokenExpired
	}

	user, err := s.user.Get(ctx, token.UserId)
	if err != nil {
		l.Errorf("Get Err:%v", err)
		return nil, err
	}

	token, err = s.generateToken(ctx, user)
	if err != nil {
		l.Error("generateToken Err:%v", err)
		return nil, err
	}

	return &v1.LoginReply{
		AccessToken:      token.AccessToken,
		ExpiresIn:        token.ExpiresIn,
		RefreshToken:     token.RefreshToken,
		RefreshExpiresIn: token.RefreshExpiresIn,
	}, nil
}

func (s *Service) generateToken(ctx context.Context, user *models.User) (*models.Token, error) {
	var (
		l   = log.Context(ctx)
		now = time.Now()
	)

	accessToken, err := s.JWT.GenerateToken(jwt.Account{
		ID:   user.ID,
		Role: int(user.Role),
	})
	if err != nil {
		l.Errorf("GenerateToken Err:%v", err)
		return nil, err
	}

	h := md5.New()
	_, err = h.Write([]byte(fmt.Sprintf("%v:%v:%v", accessToken, user.ID, time.Now().Unix())))
	if err != nil {
		l.Errorf("h.Write Err:%v", err)
		return nil, err
	}
	refreshToken := hex.EncodeToString(h.Sum(nil))

	newToken := &models.Token{
		UserId:           user.ID,
		AccessToken:      accessToken,
		ExpiresIn:        now.Unix() + s.business.JWT.TokenExpiresSeconds,
		RefreshToken:     refreshToken,
		RefreshExpiresIn: now.Unix() + s.business.JWT.RefreshExpiresSeconds,
	}

	err = s.token.Upsert(ctx, newToken)
	if err != nil {
		l.Error("token.Upsert Err:%v", err)
		return nil, err
	}
	return newToken, nil
}
