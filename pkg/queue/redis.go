package queue

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

type PublishWithRedis struct {
	client *redis.Client
}

func NewPublishWithRedis(client *redis.Client) Publisher {
	return &PublishWithRedis{
		client: client,
	}
}

func (s *PublishWithRedis) Publish(ctx context.Context, topic string, message interface{}) {
	s.client.Publish(ctx, topic, message)
}

type ConsumerWithRedis struct {
	client *redis.Client
	sub    *redis.PubSub
}

func NewConsumerWithRedis(client *redis.Client) Consumer {
	return &ConsumerWithRedis{
		client: client,
	}
}

func (s *ConsumerWithRedis) Consumer(ctx context.Context, callback Callback, topics ...string) {
	if s.sub == nil {
		s.sub = s.client.Subscribe(ctx, topics...)
	}
	ch := s.sub.Channel()
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Context(ctx).Errorf("Recovery:panic:%v", err)
			}
		}()

		for message := range ch {
			callback(message.Channel, message.Payload)
		}
	}()
}

func (s *ConsumerWithRedis) Close() {
	s.sub.Close()
}
