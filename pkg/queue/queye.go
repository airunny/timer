package queue

import "context"

type Callback func(topic, message string)

type Publisher interface {
	Publish(ctx context.Context, topic string, message interface{})
}

type Consumer interface {
	Consumer(ctx context.Context, callback Callback, topics ...string)
	Close()
}
