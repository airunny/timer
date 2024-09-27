package ikafka

import (
	"fmt"
	"testing"
	"time"

	"github.com/IBM/sarama"
	"github.com/stretchr/testify/assert"
)

func TestNewConsumer(t *testing.T) {
	consumer, err := NewConsumer(&ConsumerConfig{
		Brokers:      []string{"127.0.0.1:9092", "127.0.0.1:9093", "127.0.0.1:9094"},
		Topics:       []string{"wikifx-akamai"},
		OffsetOldest: true,
		GroupId:      "wiki_netwon",
		Version:      "2.8.1",
	})
	assert.Nil(t, err)

	err = consumer.Start(ConsumerFunc)
	assert.Nil(t, nil)
	time.Sleep(time.Minute * 10)
	consumer.Close()
}

func ConsumerFunc(session sarama.ConsumerGroupSession, message *sarama.ConsumerMessage) {
	defer session.Commit()

	fmt.Printf("Headers:%v\n Timestamp:%v\n BlockTimestamp:%v\n key:%s\n value:%s\n topic:%s\n partition:%d\n offset:%d\n",
		message.Headers,
		message.Timestamp,
		message.BlockTimestamp,
		string(message.Key),
		string(message.Value),
		message.Topic,
		message.Partition,
		message.Offset,
	)
}
