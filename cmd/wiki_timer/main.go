package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/airunny/timer/internal/conf"
	"github.com/airunny/wiki-go-tools/config"
	"github.com/airunny/wiki-go-tools/env"
	"github.com/airunny/wiki-go-tools/ilog"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	_ "go.uber.org/automaxprocs"
)

var (
	Name     = env.GetServiceName()
	Version  string
	flagconf string
	id, _    = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	opts := []kratos.Option{
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	}

	//_, ok := os.LookupEnv("KUBERNETES_SERVICE_HOST")
	//if ok {
	//	clientSet, err := k8s.NewClient()
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	opts = append(opts, kratos.Registrar(registry.NewRegistry(clientSet, logger)))
	//}

	return kratos.New(opts...)
}

func main() {
	flag.Parse()
	logger, closer := ilog.NewLogger(id, Name)
	defer closer.Close()

	var bc conf.Bootstrap
	_, err := config.LoadConfig(&bc, config.WithFilePath(flagconf))
	if err != nil {
		panic(err)
	}

	bcStr, _ := json.Marshal(&bc)
	_ = logger.Log(log.LevelInfo, "config:", string(bcStr))

	app, cleanup, err := wireApp(bc.Server, bc.Data, bc.Business, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err = app.Run(); err != nil {
		panic(err)
	}
}
