package ikafka

import (
	"fmt"
	"testing"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/stretchr/testify/assert"
)

func ConfluentConsumerFunc(c *kafka.Consumer, msg *kafka.Message) {
	defer c.CommitMessage(msg)

	fmt.Println("数据：", string(msg.Value))
}

func TestNewConfluentConsumer(t *testing.T) {
	t.SkipNow()
	consumer, err := NewConfluentConsumer(&ConsumerConfig{
		Brokers:      []string{"testkfk.fxeyeinterface.com:9092"},
		Topics:       []string{"user_behavior_data_pipline"},
		GroupId:      "wiki_netwon11",
		OffsetOldest: true,
	})
	assert.Nil(t, err)
	defer consumer.Close()

	consumer.Start(ConfluentConsumerFunc)
	assert.Nil(t, nil)

	time.Sleep(time.Minute * 1)
	return
}
