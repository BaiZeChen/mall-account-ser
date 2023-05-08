package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"mall-account-ser/account/internal/pkg"
	pkg2 "mall-account-ser/account/pkg"
	"runtime/debug"
	"strings"
)

// 权限白名单
var whiteList = map[string]struct{}{
	"Login": {},
}

// RecoveryInterceptor panic捕获
func RecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "Panic err: %v", e)
		}
	}()
	return handler(ctx, req)
}

func Auth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	arr := strings.Split(info.FullMethod, "/")
	method := arr[len(arr)-1]
	if _, ok := whiteList[method]; !ok {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "获取校验信息失败")
		}
		if token, ok := md["token"]; ok {
			jwt := &pkg.JWTClaims{}
			check, err := jwt.Check(token[0])
			if err != nil || !check {
				return nil, status.Errorf(codes.Unauthenticated, "校验身份信息失败")
			}
		} else {
			return nil, status.Errorf(codes.Unauthenticated, "获取校验信息失败")
		}
	}
	return handler(ctx, req)
}

func Limit(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	if !pkg2.GlobalLimiter.Allow() {
		return nil, status.Errorf(codes.Unavailable, "访问频繁，请稍后再试~~")
	}
	return handler(ctx, req)
}
