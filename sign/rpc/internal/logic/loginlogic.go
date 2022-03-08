package logic

import (
	"context"
	"strconv"

	"mallxx_server/common/merrorx"
	"mallxx_server/sign/rpc/internal/svc"
	"mallxx_server/sign/rpc/pb"
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

func (l *LoginLogic) Login(in *pb.MemberLoginRequest) (*pb.MemberResponse, error) {
	member := l.svcCtx.MemberModel.FindOneByUsername(in.Username)

	if member == nil {
		return nil, merrorx.NewCodeError(500, "账户或密码错误")
	}

	if utils.MD5(in.Password) != member.Password {
		return nil, merrorx.NewCodeError(500, "密码错误")
	}

	if member.Status != 0 {
		return nil, merrorx.NewCodeError(500, "账户被禁用")
	}

	token, err := utils.CreateJwtToken(member.Id, l.svcCtx.Config.JwtAuth.AccessExpire, l.svcCtx.Config.JwtAuth.AccessSecret)
	if err != nil {
		return nil, merrorx.NewCodeError(500, err.Error())
	}

	err = l.svcCtx.RedisClient.Setex(l.svcCtx.Config.TokenPrefix+token, strconv.FormatInt(member.Id, 10), int(l.svcCtx.Config.JwtAuth.AccessExpire))

	if err != nil {
		return nil, merrorx.NewCodeError(500, "登录失败")
	}

	member.Token = token
	data := &pb.MemberData{
		Member:      member,
		MemberLevel: l.svcCtx.MemberLevelModel.FindOneById(member.MemberLevel),
	}

	return &pb.MemberResponse{
		Code:   200,
		Detail: "ok",
		Data:   data,
	}, nil
}
