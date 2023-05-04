package main

import (
	"github.com/BaiZeChen/mall-api/proto/account"
	"google.golang.org/grpc"
	"mall-account-ser/account/configs"
	"mall-account-ser/account/internal/service"
	"mall-account-ser/account/pkg"
	"net"
)

func init() {
	pkg.InitGorm(configs.Conf.MySQL)
}

func main() {
	server := grpc.NewServer()
	account.RegisterAccountServiceServer(server, &service.AccountApi{})
	listen, err := net.Listen("tcp", ":"+configs.Conf.App.Port)
	if err != nil {
		panic(any(err))
	}
	err = server.Serve(listen)
	if err != nil {
		panic(any(err))
	}
}
