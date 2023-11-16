package rocket

import (
	mq "github.com/apache/rocketmq-clients/golang"
	"time"
)

// func InitLog() {
//	os.Setenv("mq.consoleAppender.enabled", "true")
//	os.Setenv("rocketmq.client.logLevel", "error")
//	mq.ResetLogger()
//}

// Consumer 消费者
type Consumer struct {
	topic string
	group string
	opt   Option
}

func NewConsumer(topic, group string, opt IOption) *Consumer {
	return &Consumer{
		topic: topic,
		group: group,
		opt:   Option{opt: opt},
	}
}

func (c *Consumer) makeOptions(await time.Duration) []mq.SimpleConsumerOption {
	return []mq.SimpleConsumerOption{
		mq.WithAwaitDuration(await),
		mq.WithSubscriptionExpressions(map[string]*mq.FilterExpression{
			c.topic: mq.SUB_ALL,
		}),
	}
}

func (c *Consumer) Connect(await time.Duration) (connect mq.SimpleConsumer) {
	var err error
	opts := c.makeOptions(await)
	config := c.opt.makeConfig()
	config.ConsumerGroup = c.group

	if connect, err = mq.NewSimpleConsumer(config, opts...); err != nil {
		panic(err)
	}

	if err = connect.Start(); err != nil {
		panic(err)
	}

	return
}
