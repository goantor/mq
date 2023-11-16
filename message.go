package rocket

import (
	mq "github.com/apache/rocketmq-clients/golang"
)

func NewMessage(body []byte) *mq.Message {
	return &mq.Message{Body: body}
}
