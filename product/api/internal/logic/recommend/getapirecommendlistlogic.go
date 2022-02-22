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

type GetApiRecommendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApiRecommendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetApiRecommendListLogic {
	return GetApiRecommendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApiRecommendListLogic) GetApiRecommendList(req types.ProductListRequest) (*types.ProductListRespone, error) {
	var rpcReq productservices.ProductListRequest
	err := copier.Copy(&rpcReq, req)
	if err != nil {
		return nil, merrorx.NewCodeError(500, err.Error())
	}

	resp, err := l.svcCtx.ProductRpc.GetApiRecommendList(l.ctx, &rpcReq)

	if err != nil {
		return nil, err
	}

	var data []*types.Product
	copier.Copy(&data, resp.Data)

	return &types.ProductListRespone{
		Code:   resp.Code,
		Detail: resp.Detail,
		Data:   data,
	}, nil
}
