package logic

import (
	"context"
	"fmt"

	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBrandInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBrandInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBrandInfoLogic {
	return &GetBrandInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBrandInfoLogic) GetBrandInfo(in *pb.BrandRequest) (*pb.BrandOneResponse, error) {
	brand := l.svcCtx.BrandModel.FindOne(in.Id)
	fmt.Println(in)
	return &pb.BrandOneResponse{
		Code:   200,
		Data:   brand,
		Detail: "ok",
	}, nil
}
