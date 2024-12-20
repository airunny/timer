package server

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/airunny/timer/api/common"
	v1 "github.com/airunny/timer/api/timer/v1"
	innErr "github.com/airunny/timer/errors"
	"github.com/airunny/timer/internal/service"
	"github.com/airunny/timer/pkg/icontext"
	"github.com/airunny/wiki-go-tools/env"
	"github.com/airunny/wiki-go-tools/ilog"
	"github.com/airunny/wiki-go-tools/metrics"
	"github.com/airunny/wiki-go-tools/middleware"
	"github.com/go-kratos/grpc-gateway/v2/protoc-gen-openapiv2/generator"
	"github.com/go-kratos/kratos/v2/log"
	kratosMid "github.com/go-kratos/kratos/v2/middleware"
	mmd "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewHTTPServer(c *common.ServerConfig, svc *service.Service, logger log.Logger) *http.Server {
	serviceName := env.GetServiceName()
	var opts = []http.ServerOption{
		http.Filter(middleware.CORS(), ilog.LoggingHandler(serviceName)),
		http.Middleware(
			mmd.Server(),
			tracing.Server(),
			recovery.Recovery(),
			middleware.TraceIdAndRequestIdWithHeader,
			metrics.ServerMetricsMiddleware(serviceName),
			validate.Validator(),
			ParseToken(svc),
		),
		http.ErrorEncoder(middleware.ErrorEncoder),
		http.ResponseEncoder(middleware.ResponseEncoder),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.TimeoutSeconds > 0 {
		opts = append(opts, http.Timeout(time.Duration(c.Http.TimeoutSeconds)*time.Second))
	}
	srv := http.NewServer(opts...)
	v1.RegisterServiceHTTPServer(srv, svc)

	openAPIHandler := openapiv2.NewHandler(openapiv2.WithGeneratorOptions(
		generator.UseJSONNamesForFields(false),
		generator.EnumsAsInts(true),
	))
	srv.HandlePrefix("/q/", openAPIHandler)
	srv.Handle("/metrics", promhttp.Handler())
	return srv
}

func ParseToken(svc *service.Service) kratosMid.Middleware {
	return func(handler kratosMid.Handler) kratosMid.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return handler(ctx, req)
			}

			// 登录接口不需要鉴权
			if tr.Operation() == v1.OperationServiceLogin {
				return handler(ctx, req)
			}

			if tr.RequestHeader().Get("Authorization") == "" {
				return nil, innErr.ErrLogin
			}

			token := strings.TrimSpace(strings.TrimPrefix(tr.RequestHeader().Get("Authorization"), "Bearer"))
			ac, err := svc.JWT.ParseToken(token)
			if err != nil && !errors.Is(err, jwt.ErrTokenExpired) {
				return nil, innErr.ErrLogin
			}

			// 到这里说明access token过期
			if err != nil {
				// 如果账号信息为空，则重新登录
				if ac == nil {
					return nil, innErr.ErrLogin
				}

				// 如果refresh token已过期，则重新登录
				if ac.RefreshTokenExpiresAt < time.Now().Unix() {
					return nil, innErr.ErrLogin
				}

				// 如果不是刷新token，则提醒需要刷新token
				if tr.Operation() != v1.OperationServiceRefreshToken {
					return nil, innErr.ErrAccessTokenExpired
				}
				return handler(ctx, req)
			}
			// 做权限校验
			err = verifyRole(tr.Operation(), v1.UserRole(ac.Role))
			if err != nil {
				return nil, err
			}
			ctx = icontext.WithUserId(ctx, ac.ID)
			ctx = icontext.WithRole(ctx, strconv.Itoa(ac.Role))
			return handler(ctx, req)
		}
	}
}

var (
	writeOption = map[string]struct{}{
		v1.OperationServiceAddApplication:          {},
		v1.OperationServiceGenApplicationSecret:    {},
		v1.OperationServiceUpdateApplicationStatus: {},
		v1.OperationServiceUpdateApplication:       {},
		v1.OperationServiceDeleteApplication:       {},
		v1.OperationServiceAddTimer:                {},
		v1.OperationServiceReplayTimer:             {},
		v1.OperationServiceDeleteTimer:             {},
	}
	readOption = map[string]struct{}{
		v1.OperationServiceGetApplication:  {},
		v1.OperationServiceGetTimer:        {},
		v1.OperationServiceListApplication: {},
		v1.OperationServiceListTimer:       {},
		v1.OperationServiceListTask:        {},
	}
)

func verifyRole(operation string, role v1.UserRole) error {
	switch role {
	case v1.UserRole_ReadWrite:
		if _, ok := writeOption[operation]; ok {
			return nil
		}
		fallthrough
	case v1.UserRole_ReadOnly:
		if _, ok := readOption[operation]; ok {
			return nil
		}
	default:
		return nil
	}
	return innErr.ErrForbidden
}
