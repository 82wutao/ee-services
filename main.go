package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/82wutao/ee-rpcdeclare/mq"
	"github.com/82wutao/ee-rpcdeclare/network"
	"github.com/82wutao/ee-rpcdeclare/rpcx"
	"github.com/82wutao/ee-services/interfaces"
	"github.com/82wutao/ee-services/services"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatalf("services listen_addr listen_port trans_proto\n")
		return
	}
	services.Loop(100)

	opt := mq.RocketMQOptions{
		NameserverAddrs: []string{"http://localhost:9876"},
		GroupName:       "newG",
		ProduceRetries:  2,
		ConsumeMode:     consumer.Clustering,
	}
	services.StartMqProducer(opt)
	services.StartMqConsumer(opt, []mq.MQSubscribe{
		mq.MQSubscribe{
			Topic: "TopicTest",
			Tag:   "",
			Callback: func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {

				for _, msg := range msgs {
					log.Printf("rcv a msg from mq %s\n", (string)(msg.Body))
				}
				return consumer.ConsumeSuccess, nil
			},
		},
	})

	addr := os.Args[1]
	port, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("get port argument error %+v\n", err)
		return
	}
	proto := os.Args[3]

	if err := interfaces.LaunchRpcServer(context.Background(), network.HostPort{
		Host:  addr,
		Port:  int16(port),
		Proto: proto,
	}, []rpcx.ServiceHandle{new(interfaces.OrderService)}); err != nil {
		log.Fatalf("LaunchRpcServer error %+v\n", err)
	}
	log.Printf("LaunchRpcServer finished \n")
}
