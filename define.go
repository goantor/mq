package rocket

import (
	mq "github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
	"time"
)

type IOption interface {
	TakeCredentials() *credentials.SessionCredentials
	TakeConfig() *mq.Config
	TakeTopic() string
	TakeGroup() string
}

type IProducerOption interface {
	IOption
	TakeOptions() []mq.ProducerOption
}

type IConsumerOption interface {
	IOption
	TakeOptions() []mq.SimpleConsumerOption
	TakeNum() int
	TakeWait() time.Duration
	TakeMaxMessageNum() int32
	TakeInvisibleDuration() time.Duration
}

type IConsumer interface {
	Connect() (connect mq.SimpleConsumer)
}

type IProducer interface {
	Connect() (connect mq.Producer)
}
