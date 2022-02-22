package logic

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBrandLogic {
	return &DeleteBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteBrandLogic) DeleteBrand(in *pb.BrandRequest) (*pb.Response, error) {
	if l.svcCtx.BrandModel.Delete(in.Id) == false {
		return nil, merrorx.NewCodeError(500, "新增失败")
	}

	return &pb.Response{
		Code:   200,
		Detail: "新增成功",
	}, nil
}
