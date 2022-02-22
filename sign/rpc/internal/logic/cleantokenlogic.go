package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/sign/rpc/internal/svc"
	"mallxx_server/sign/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CleanTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCleanTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CleanTokenLogic {
	return &CleanTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CleanTokenLogic) CleanToken(in *pb.SignTokenRequest) (*pb.Response, error) {
	if _, err := l.svcCtx.RedisClient.Del(l.svcCtx.Config.TokenPrefix + in.Token); err != nil {
		return nil, merrorx.NewCodeError(500, "Delete token is Failed")
	}
	return &pb.Response{
		Code:   200,
		Detail: "ok",
	}, nil
}
