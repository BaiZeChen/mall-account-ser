package main

import (
	"github.com/BaiZeChen/mall-api/proto/account"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
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
	pkg.InitTracing()
}

func main() {
	opst := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			interceptor.RecoveryInterceptor,
			otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer(), otgrpc.LogPayloads()),
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
	err = pkg.TracingCloser.Close()
	if err != nil {
		panic(err)
	}
}
