package logic

import (
	"context"
	"strconv"

	"mallxx_server/common/merrorx"
	"mallxx_server/sign/rpc/internal/svc"
	"mallxx_server/sign/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyTokenLogic {
	return &VerifyTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyTokenLogic) VerifyToken(in *pb.SignTokenRequest) (*pb.Response, error) {
	uid, err := l.svcCtx.RedisClient.Get(l.svcCtx.Config.TokenPrefix + in.Token)
	if err != nil {
		return nil, merrorx.NewCodeError(500, "Token is not found")
	}

	if uid != strconv.FormatInt(in.Id, 10) {
		return nil, merrorx.NewCodeError(401, "Verify totken is Failed")
	}
	return &pb.Response{
		Code:   200,
		Detail: "ok",
	}, nil
}
