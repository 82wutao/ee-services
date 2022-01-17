package interfaces

import (
	"context"
	"fmt"
	"log"

	"github.com/smallnest/rpcx/server"
)

func _onRestart(s *server.Server)  {}
func _onShutdown(s *server.Server) {}

type HostPort struct {
	Host  string
	Port  int16
	Proto string // tcp/udp/http
}

func (hp HostPort) toString() string {
	return fmt.Sprintf("%s@%s:%d",
		hp.Proto, hp.Host, hp.Port)
}

type ServiceHandle interface {
	HandleName() string
}

var serv *server.Server

func LaunchRpcServer(serviceHost HostPort, handles []ServiceHandle) error {

	if serv != nil {
		return nil
	}

	serv = server.NewServer()

	// rp := serverplugin.ConsulRegisterPlugin{
	// 	ServiceAddress: serviceHost.toString(),
	// 	ConsulServers:  []string{""},
	// 	BasePath:       "ee/rpc",
	// 	Metrics:        metrics.NewRegistry(),
	// 	UpdateInterval: time.Minute,
	// }
	// if err := rp.Start(); err != nil {
	// 	log.Fatalf("regist service error %+v\n", err)
	// 	return err
	// }
	// serv.Plugins.Add(rp) //consul
	// serv.Plugins.Add(nil) //trace

	serv.RegisterOnRestart(_onRestart)   // on restart
	serv.RegisterOnShutdown(_onShutdown) // on shutdown

	for _, hanle := range handles {
		if err := serv.RegisterName(hanle.HandleName(), hanle, ""); err != nil {
			log.Fatalf("server regist service %s error %+v\n", hanle.HandleName(), err)
			return err
		}
	}
	if err := serv.Serve(serviceHost.Proto,
		fmt.Sprintf("%s:%d", serviceHost.Host, serviceHost.Port)); err != nil {
		log.Fatalf("start rpcx server error %+v\n", err)
		return err
	}
	return nil
}
func ShutdownRpcServer(ctx context.Context) error {
	if serv == nil {
		return nil
	}
	if err := serv.Shutdown(ctx); err != nil {
		log.Fatalf("shutdown rpcx server error %+v\n", err)
		return err
	}
	return nil
}
func RelauchRpcServer(ctx context.Context) error {
	if serv == nil {
		return nil
	}
	if err := serv.Restart(ctx); err != nil {
		log.Fatalf("restart rpcx server error %+v\n", err)
		return err
	}
	return nil
}
