package server

import (
	"time"

	"github.com/airunny/timer/api/common"
	v1 "github.com/airunny/timer/api/timer/v1"
	"github.com/airunny/timer/internal/service"
	"github.com/airunny/wiki-go-tools/env"
	"github.com/airunny/wiki-go-tools/ilog"
	"github.com/airunny/wiki-go-tools/metrics"
	"github.com/airunny/wiki-go-tools/middleware"
	"github.com/go-kratos/grpc-gateway/v2/protoc-gen-openapiv2/generator"
	"github.com/go-kratos/kratos/v2/log"
	mmd "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
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
