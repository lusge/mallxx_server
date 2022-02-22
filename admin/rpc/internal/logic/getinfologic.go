package logic

import (
	"context"
	"fmt"

	"mallxx_server/admin/rpc/internal/svc"
	"mallxx_server/admin/rpc/pb"
	"mallxx_server/common/merrorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfoLogic {
	return &GetInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetInfoLogic) GetInfo(in *pb.AdminInfo) (*pb.AdminInfoResponse, error) {
	// todo: add your logic here and delete this line
	adminInfo := l.svcCtx.AdminModel.FindOne(in.Id)
	fmt.Println(in.Id)
	if adminInfo == nil {
		return nil, merrorx.NewCodeError(500, "Admin not found")
	}

	return &pb.AdminInfoResponse{
		Code:   200,
		Data:   adminInfo,
		Detail: "ok",
	}, nil
}
