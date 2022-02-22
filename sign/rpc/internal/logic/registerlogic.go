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

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.Member) (*pb.MemberResponse, error) {
	member := l.svcCtx.MemberModel.FindOneByUsername(in.Username)

	if member != nil {
		return nil, merrorx.NewCodeError(500, "该用户被注册")
	}

	isOk, info := l.svcCtx.MemberModel.Register(in.Username, utils.MD5(in.Password))
	if isOk == false {
		return nil, merrorx.NewCodeError(500, "注册失败")
	}

	token, err := utils.CreateJwtToken(member.Id, l.svcCtx.Config.JwtAuth.AccessExpire, l.svcCtx.Config.JwtAuth.AccessSecret)
	if err != nil {
		return nil, merrorx.NewCodeError(500, err.Error())
	}

	err = l.svcCtx.RedisClient.Setex(l.svcCtx.Config.TokenPrefix+token, strconv.FormatInt(member.Id, 10), int(l.svcCtx.Config.JwtAuth.AccessExpire))

	if err != nil {
		return nil, merrorx.NewCodeError(500, "登录失败")
	}

	data := &pb.MemberData{
		Member:      info,
		MemberLevel: l.svcCtx.MemberLevelModel.FindOneById(info.MemberLevel),
	}

	return &pb.MemberResponse{
		Code:   200,
		Detail: "ok",
		Data:   data,
	}, nil
}
