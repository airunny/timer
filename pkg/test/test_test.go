package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/airunny/timer/api/common"
	"github.com/airunny/timer/internal/dao"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/zeebo/assert"
)

func TestRedis(t *testing.T) {
	redisClient, closer, err := dao.NewRedis(&common.DataConfig{
		Redis: &common.DataConfig_Redis{
			Address: "127.0.0.1:6379",
		},
	}, log.GetLogger())
	assert.Nil(t, err)
	defer closer()

	sub := redisClient.Subscribe(context.Background(), "new")
	ch := sub.Channel()
	go func() {
		for value := range ch {
			fmt.Println("message", value.Channel, value.Pattern, value.Payload, value.PayloadSlice)
		}
	}()

	for i := 0; i <= 100; i++ {
		redisClient.Publish(context.Background(), "new", fmt.Sprintf("test:%d", i))
	}
}
