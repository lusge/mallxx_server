package brand

import (
	"context"
	"fmt"

	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
	"mallxx_server/product/rpc/productservices"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetBrandInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBrandInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetBrandInfoLogic {
	return GetBrandInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBrandInfoLogic) GetBrandInfo(req types.BrandRequest) (*types.BrandOneResponse, error) {
	fmt.Println(req, "....")
	resp, err := l.svcCtx.ProductRpc.GetBrandInfo(l.ctx, &productservices.BrandRequest{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Keyword:  req.Keyword,
		Id:       req.Id,
	})

	if err != nil {
		return nil, err
	}

	var data types.Brand
	copier.Copy(&data, resp.Data)

	return &types.BrandOneResponse{
		Code:   resp.Code,
		Detail: resp.Detail,
		Data:   &data,
	}, nil
}
