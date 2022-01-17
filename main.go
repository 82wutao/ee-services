package main

import (
	"log"
	"os"
	"strconv"

	"github.com/82wutao/ee-services/interfaces"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatalf("services listen_addr listen_port trans_proto\n")
		return
	}
	addr := os.Args[1]
	port, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("get port argument error %+v\n", err)
		return
	}
	proto := os.Args[3]

	if err := interfaces.LaunchRpcServer(interfaces.HostPort{
		Host:  addr,
		Port:  int16(port),
		Proto: proto, // tcp/udp/http
	}, []interfaces.ServiceHandle{new(interfaces.OrderService)}); err != nil {
		log.Fatalf("LaunchRpcServer error %+v\n", err)
	}
}
