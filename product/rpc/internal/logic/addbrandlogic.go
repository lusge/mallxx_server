package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddBrandLogic {
	return &AddBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddBrandLogic) AddBrand(in *pb.Brand) (*pb.Response, error) {
	isOk := l.svcCtx.BrandModel.Insert(in)
	if !isOk {
		return nil, merrorx.NewCodeError(500, "新增失败")
	}

	return &pb.Response{
		Code:   200,
		Detail: "ok",
	}, nil
}
