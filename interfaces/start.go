package interfaces

import (
	"context"
	"log"

	"github.com/82wutao/ee-rpcdeclare/rpcx"
	"github.com/smallnest/rpcx/server"
)

func _onRestart(s *server.Server)  {}
func _onShutdown(s *server.Server) {}

var serv *rpcx.RPCXServer

func LaunchRpcServer(ctx context.Context, serviceHost rpcx.HostPort, handles []rpcx.ServiceHandle) error {

	if serv != nil {
		return nil
	}

	var err error
	serv, err = rpcx.NewServerAndRegisterConsul(serviceHost, handles,
		_onRestart, _onShutdown,
		rpcx.HostPort{Proto: "http", Host: "127.0.0.1", Port: 8500})
	// serv, err = rpcx.NewServer()
	if err != nil {
		return err
	}
	return serv.Launch(ctx)
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
	if err := serv.Relaunch(ctx); err != nil {
		log.Fatalf("restart rpcx server error %+v\n", err)
		return err
	}
	return nil
}
