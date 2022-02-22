package logic

import (
	"context"

	"mallxx_server/product/rpc/internal/svc"
	"mallxx_server/product/rpc/pb"
	"mallxx_server/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductAttributeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductAttributeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductAttributeListLogic {
	return &GetProductAttributeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductAttribute
func (l *GetProductAttributeListLogic) GetProductAttributeList(in *pb.AttrRequest) (*pb.ProductAttributeListResponse, error) {
	start, size := utils.Pager(in.PageNum, in.PageSize)

	productAttrModel := l.svcCtx.ProductAttributeModel

	list := productAttrModel.FindByPage(in.Cid, in.Keyword, in.Type, start, size)
	total := productAttrModel.Count(in.Cid, in.Keyword, in.Type)

	return &pb.ProductAttributeListResponse{
		Code:   200,
		Detail: "ok",
		Total:  total,
		Data:   list,
	}, nil
}
