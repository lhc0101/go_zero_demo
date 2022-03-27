package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"mall/common/js2session"
	"mall/service/wxuser/model"
	"mall/service/wxuser/rpc/internal/svc"
	"mall/service/wxuser/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAppLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppLoginLogic {
	return &AppLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AppLoginLogic) AppLogin(in *user.AppLoginRequest) (*user.AppLoginResponse, error) {
	// todo: add your logic here and delete this line
	userSession, err := js2session.Code2session(in.Code, in.AppId, js2session.WX_APP_SECRET)
	res, err := l.svcCtx.UserModel.FindOneByOpenId(userSession.Openid)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	return &user.AppLoginResponse{
		Id:       res.Id,
		Name:     res.Name,
		Gender:   res.Gender,
		Mobile:   res.Mobile,
		OpenId:   res.OpenId,
		Nickname: res.NickName,
	}, nil

}
