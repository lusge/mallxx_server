package logic

import (
	"context"
	"fmt"
	"strconv"

	"mallxx_server/admin/rpc/internal/svc"
	"mallxx_server/admin/rpc/pb"
	"mallxx_server/common/merrorx"

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

func (l *VerifyTokenLogic) VerifyToken(in *pb.AdminTokenRequest) (*pb.Response, error) {
	// todo: add your logic here and delete this line
	uid, err := l.svcCtx.RedisClient.Get(in.Token)
	fmt.Println(in.Token, err)
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
