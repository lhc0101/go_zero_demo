// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package userclient

import (
	"context"

	"mall/service/wxuser/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AppLoginRequest  = user.AppLoginRequest
	AppLoginResponse = user.AppLoginResponse
	UserInfoRequest  = user.UserInfoRequest
	UserInfoResponse = user.UserInfoResponse

	User interface {
		AppLogin(ctx context.Context, in *AppLoginRequest, opts ...grpc.CallOption) (*AppLoginResponse, error)
		UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) AppLogin(ctx context.Context, in *AppLoginRequest, opts ...grpc.CallOption) (*AppLoginResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.AppLogin(ctx, in, opts...)
}

func (m *defaultUser) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}
