package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"mall/service/wxuser/api/internal/svc"
	"mall/service/wxuser/api/internal/types"
	"mall/service/wxuser/rpc/userclient"
)

type AppLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAppLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) AppLoginLogic {
	return AppLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AppLoginLogic) AppLogin(req types.AppLoginRequest) (resp *types.AppLoginResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.UserRpc.AppLogin(l.ctx, &userclient.AppLoginRequest{
		AppId:    req.AppId,
		Code:     req.Code,
		Nickname: req.Nickname,
	})
	if err != nil {
		return nil, err
	}

	//now := time.Now().Unix()
	//accessExpire := l.svcCtx.Config.Auth.AccessExpire
	//
	//accessToken, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, res.Id)
	//if err != nil {
	//	return nil, err
	//}

	return &types.AppLoginResponse{
		Message:     "user authorize successfully.",
		Result_code: 0,
	}, nil

}
