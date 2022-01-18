package services

import (
	"context"

	"github.com/82wutao/ee-rpcdeclare/mq"
)

var p *mq.MQProducer
var c *mq.MQConsumer

func StartMqProducer(opt mq.RocketMQOptions) error {
	var err error
	p, err = mq.NewRocketMQProducer(opt)
	if err != nil {
		return err
	}

	return p.Start()
}
func StopMqProducer() error {
	if p == nil {
		return nil
	}
	return p.Shutdown()
}

func SendMsg(ctx context.Context, msg mq.MQSending) (bool, error) {
	return p.Sync(ctx, msg)
}

func StartMqConsumer(opt mq.RocketMQOptions, subscribes []mq.MQSubscribe) error {

	_c, _err := mq.NewRocketMQConsumer(opt)
	if _err != nil {
		return _err
	}
	for _, s := range subscribes {
		if err := _c.Subscribe(s); err != nil {
			return err
		}
	}
	c = _c
	return _c.Start()
}
func StopMqConsumer() error {
	if c == nil {
		return nil
	}
	return c.Shutdown()
}
