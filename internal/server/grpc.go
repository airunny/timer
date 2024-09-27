package server

import (
	"time"

	"github.com/airunny/timer/api/common"
	v1 "github.com/airunny/timer/api/timer/v1"
	"github.com/airunny/timer/internal/service"
	"github.com/airunny/wiki-go-tools/env"
	"github.com/airunny/wiki-go-tools/ilog"
	"github.com/airunny/wiki-go-tools/metrics"
	"github.com/go-kratos/kratos/v2/log"
	mmd "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewGRPCServer(c *common.ServerConfig, svc *service.Service, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			mmd.Server(),
			tracing.Server(),
			recovery.Recovery(),
			ilog.LoggingGRPC(),
			metrics.ServerMetricsMiddleware(env.GetServiceName()),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.TimeoutSeconds > 0 {
		opts = append(opts, grpc.Timeout(time.Duration(c.Grpc.TimeoutSeconds)*time.Second))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterServiceServer(srv, svc)
	return srv
}
