package service

import (
	"context"
	"github.com/BaiZeChen/mall-api/proto/account"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"mall-account-ser/account/internal/biz"
)

type AccountApi struct{}

func (a *AccountApi) CreateAccount(ctx context.Context, req *account.ReqAddAccount) (*emptypb.Empty, error) {
	control := &biz.AccountControl{
		Name:     req.Name,
		Password: req.Password,
	}
	if len(control.Name) == 0 || len(control.Name) > 12 {
		return nil, status.Errorf(codes.InvalidArgument, "账号长度不符合规则，请重新输入！")
	}
	if len(control.Password) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "请填写密码！")
	}
	err := control.Add()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "添加账户失败，原因：%s", err.Error())
	}
	return &emptypb.Empty{}, nil
}
func (a *AccountApi) UpdateAccountName(ctx context.Context, req *account.ReqUpdateAccountName) (*emptypb.Empty, error) {
	return nil, nil
}
func (a *AccountApi) UpdateAccountPassword(context.Context, *account.ReqUpdateAccountPassword) (*emptypb.Empty, error) {
	return nil, nil
}
func (a *AccountApi) DeleteAccount(context.Context, *account.ReqDelAccount) (*emptypb.Empty, error) {
	return nil, nil
}
func (a *AccountApi) AccountList(context.Context, *account.ReqAccountList) (*account.List, error) {
	return nil, nil
}
