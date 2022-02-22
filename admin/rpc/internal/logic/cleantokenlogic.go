package logic

import (
	"context"

	"mallxx_server/admin/rpc/internal/svc"
	"mallxx_server/admin/rpc/pb"
	"mallxx_server/common/merrorx"

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

func (l *CleanTokenLogic) CleanToken(in *pb.AdminTokenRequest) (*pb.Response, error) {
	// todo: add your logic here and delete this line
	if _, err := l.svcCtx.RedisClient.Del(in.Token); err != nil {
		return nil, merrorx.NewCodeError(500, "Delete token is Failed")
	}
	return &pb.Response{
		Code:   200,
		Detail: "ok",
	}, nil
}
