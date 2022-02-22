package logic

import (
	"context"
	"fmt"
	"strconv"

	"mallxx_server/admin/rpc/internal/svc"
	"mallxx_server/admin/rpc/pb"
	"mallxx_server/common/merrorx"
	"mallxx_server/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.AdminInfo) (*pb.AdminInfoResponse, error) {
	// todo: add your logic here and delete this line
	adminInfo := l.svcCtx.AdminModel.FindByUsername(in.Username)
	if adminInfo == nil {
		return nil, merrorx.NewCodeError(500, "用户不存在")
	}

	if adminInfo.Password != utils.MD5(in.Password) {
		fmt.Println(adminInfo)
		return nil, merrorx.NewCodeError(500, "密码错误")
	}

	adminInfo.Password = ""

	token, err := utils.CreateJwtToken(adminInfo.Id, l.svcCtx.Config.JwtAuth.AccessExpire, l.svcCtx.Config.JwtAuth.AccessSecret)
	if err != nil {
		return nil, merrorx.NewCodeError(500, err.Error())
	}

	err = l.svcCtx.RedisClient.Setex(l.svcCtx.Config.TokenPrefix+token, strconv.FormatInt(adminInfo.Id, 10), int(l.svcCtx.Config.JwtAuth.AccessExpire))

	if err != nil {
		return nil, merrorx.NewCodeError(500, "登录失败")
	}

	adminInfo.Token = token
	return &pb.AdminInfoResponse{
		Code:   200,
		Detail: "ok",
		Data:   adminInfo,
	}, nil
}
