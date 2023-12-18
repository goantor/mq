package rocket

import (
	"fmt"
	mq "github.com/apache/rocketmq-clients/golang"
)

// func InitLog() {
//	os.Setenv("mq.consoleAppender.enabled", "true")
//	os.Setenv("rocketmq.client.logLevel", "error")
//	mq.ResetLogger()
//}

// Consumer 消费者
type Consumer struct {
	opt IConsumerOption
}

func NewConsumer(opt IConsumerOption) *Consumer {
	return &Consumer{
		opt: opt,
	}
}

func (c *Consumer) Connect() (connect mq.SimpleConsumer) {
	var err error
	//fmt.Printf()
	fmt.Printf("%v\n", c.opt)
	if connect, err = mq.NewSimpleConsumer(c.opt.TakeConfig(), c.opt.TakeOptions()...); err != nil {
		panic(err)
	}

	if err = connect.Start(); err != nil {
		panic(err)
	}

	return
}
