package rocket

import (
	mq "github.com/apache/rocketmq-clients/golang"
)

type Producer struct {
	topic string
	group string
	opt   Option
}

func NewProducer(topic string, opt IOption) IProducer {
	return &Producer{
		opt: Option{
			opt: opt,
		},
		topic: topic,
	}
}

func (c *Producer) makeOptions() []mq.ProducerOption {
	return []mq.ProducerOption{
		mq.WithTopics(c.topic),
	}
}

func (c *Producer) Connect() (product mq.Producer) {
	var err error
	config := c.opt.makeConfig()
	config.ConsumerGroup = c.group

	if product, err = mq.NewProducer(config, c.makeOptions()...); err != nil {
		panic(err)
	}

	if err = product.Start(); err != nil {
		panic(err)
	}

	return
}
