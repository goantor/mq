package rocket

import (
	mq "github.com/apache/rocketmq-clients/golang"
)

type Producer struct {
	opt IProducerOption
}

func NewProducer(opt IProducerOption) IProducer {
	return &Producer{
		opt: opt,
	}
}

//func (c *Producer) makeOptions() []mq.ProducerOption {
//	return []mq.ProducerOption{
//		mq.WithTopics(c.opt.TakeTopic()),
//	}
//}

func (c *Producer) Connect() (product mq.Producer) {
	var err error
	if product, err = mq.NewProducer(c.opt.TakeConfig(), c.opt.TakeOptions()...); err != nil {
		panic(err)
	}

	if err = product.Start(); err != nil {
		panic(err)
	}

	return
}
