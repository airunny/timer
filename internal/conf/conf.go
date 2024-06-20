package conf

import "timer/api/common"

type Bootstrap struct {
	Server   *common.ServerConfig `json:"server"`
	Data     *common.DataConfig   `json:"data"`
	Business *Business            `json:"business"`
}

type Elastic struct {
	Source   string `json:"source"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Business struct {
	Elastic *Elastic `json:"elastic"`
}
