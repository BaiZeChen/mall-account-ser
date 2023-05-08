package main

import (
	"github.com/BaiZeChen/mall-api/proto/account"
	"google.golang.org/grpc"
	"mall-ser/account/configs"
	"mall-ser/account/internal/interceptor"
	"mall-ser/account/internal/service"
	"mall-ser/account/pkg"
	"net"
)

func init() {
	pkg.InitGorm(configs.Conf.MySQL)
	pkg.FlowControl()
}

func main() {
	opst := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			interceptor.RecoveryInterceptor,
			interceptor.Limit,
			interceptor.Auth,
		),
	}
	server := grpc.NewServer(opst...)
	account.RegisterAccountServiceServer(server, &service.AccountApi{})
	listen, err := net.Listen("tcp", ":"+configs.Conf.App.Port)
	if err != nil {
		panic(err)
	}
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
