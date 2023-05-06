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

func (a *AccountApi) Login(ctx context.Context, req *account.ReqAddAccount) (*account.RespToken, error) {
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

	token, err := control.Login()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "登录失败，原因：%s", err.Error())
	}
	return &account.RespToken{Token: token}, nil
}
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
	control := &biz.AccountControl{
		ID:   uint(req.Id),
		Name: req.Name,
	}
	if control.ID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "没有找到相对应的用户")
	}
	if len(control.Name) == 0 || len(control.Name) > 12 {
		return nil, status.Errorf(codes.InvalidArgument, "账号长度不符合规则，请重新输入！")
	}
	err := control.UpdateName()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "修改账户名称失败，原因：%s", err.Error())
	}
	return &emptypb.Empty{}, nil
}
func (a *AccountApi) UpdateAccountPassword(ctx context.Context, req *account.ReqUpdateAccountPassword) (*emptypb.Empty, error) {
	control := &biz.AccountControl{
		ID:       uint(req.Id),
		Password: req.Password,
	}
	if control.ID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "没有找到相对应的用户")
	}
	if len(control.Password) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "请填写密码！")
	}
	err := control.UpdatePassword()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "修改密码失败，原因：%s", err.Error())
	}
	return &emptypb.Empty{}, nil
}
func (a *AccountApi) DeleteAccount(ctx context.Context, req *account.ReqDelAccount) (*emptypb.Empty, error) {
	control := &biz.AccountControl{
		ID: uint(req.Id),
	}
	if control.ID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "没有找到相对应的用户")
	}
	err := control.Delete()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "删除账户失败，原因：%s", err.Error())
	}
	return &emptypb.Empty{}, nil
}
func (a *AccountApi) AccountList(ctx context.Context, req *account.ReqAccountList) (*account.List, error) {
	offset := req.Offset
	limit := req.Limit
	if offset < 0 || limit <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "页码参数不对")
	}
	if limit > 50 {
		limit = 50
	}
	control := &biz.AccountControl{
		Name: req.Name,
	}
	list, count, err := control.List(int(offset), int(limit))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取账户列表失败，原因：%s", err.Error())
	}

	result := &account.List{}
	result.Count = uint32(count)
	for _, value := range list {
		result.Data = append(result.Data, &account.Account{
			Id:         uint32(value.ID),
			Name:       value.Name,
			CreateTime: uint64(value.CreateTime),
			UpdateTime: uint64(value.UpdateTime),
		})
	}
	return result, nil
}
