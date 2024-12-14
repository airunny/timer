package conf

import (
	"github.com/airunny/timer/api/common"
	"github.com/airunny/timer/pkg/ikafka"
	"github.com/airunny/timer/pkg/jwt"
)

type Bootstrap struct {
	Server   *common.ServerConfig `json:"server"`
	Data     *common.DataConfig   `json:"data"`
	Business *Business            `json:"business"`
}

type Business struct {
	KAFKA *ikafka.ProducerConfig `json:"kafka"`
	JWT   *jwt.Config            `json:"jwt"`
}
