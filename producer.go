package rocket

import (
	mq "github.com/apache/rocketmq-clients/golang"
)

type Producer struct {
	*Option
}

func NewProducer(opt IOption) IProducer {
	return &Producer{
		&Option{
			opt: opt,
		},
	}
}

func (c *Producer) makeOptions() []mq.ProducerOption {
	return []mq.ProducerOption{
		mq.WithTopics(c.opt.TakeTopic()),
	}
}

func (c *Producer) Connect() (product mq.Producer) {
	var err error
	config := c.makeConfig()
	config.ConsumerGroup = c.opt.TakeGroup()

	if product, err = mq.NewProducer(config, c.makeOptions()...); err != nil {
		panic(err)
	}

	if err = product.Start(); err != nil {
		panic(err)
	}

	return
}
