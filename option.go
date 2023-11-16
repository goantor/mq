package rocket

import (
	mq "github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
)

type Option struct {
	opt IOption
}

func (o *Option) makeCredentials() *credentials.SessionCredentials {
	return &credentials.SessionCredentials{
		AccessKey:    o.opt.TakeAccessKey(),
		AccessSecret: o.opt.TakeSecretKey(),
	}
}

func (o *Option) makeConfig() *mq.Config {
	return &mq.Config{
		Endpoint:    o.opt.TakeEndpoint(),
		Credentials: o.makeCredentials(),
	}
}
