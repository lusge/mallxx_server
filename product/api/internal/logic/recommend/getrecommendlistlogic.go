package recommend

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetRecommendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRecommendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetRecommendListLogic {
	return GetRecommendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRecommendListLogic) GetRecommendList(req types.ProductListRequest) (*types.RecommendListResponse, error) {
	var rpcReq productservices.ProductListRequest
	err := copier.Copy(&rpcReq, req)
	if err != nil {
		return nil, merrorx.NewCodeError(500, err.Error())
	}

	resp, err := l.svcCtx.ProductRpc.GetRecommendList(l.ctx, &rpcReq)

	if err != nil {
		return nil, err
	}

	var data []*types.ProductRecommend
	copier.Copy(&data, resp.Data)

	return &types.RecommendListResponse{
		Code:   resp.Code,
		Detail: resp.Detail,
		Data:   data,
	}, nil
}
