package rocket

import (
	mq "github.com/apache/rocketmq-clients/golang"
	"time"
)

type IOption interface {
	TakeEndpoint() string
	TakeAccessKey() string
	TakeSecretKey() string
	IsDebug() string
	DebugLevel() string
}

type IConsumer interface {
	Connect(await time.Duration) (connect mq.SimpleConsumer)
}

type IProducer interface {
	Connect() (connect mq.Producer)
}
