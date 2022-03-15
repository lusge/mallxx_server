package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"mallxx_server/cart/rpc/internal/svc"
	"mallxx_server/cart/rpc/pb"
	"mallxx_server/common/merrorx"
	"mallxx_server/common/redisx"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModiryCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModiryCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModiryCartLogic {
	return &ModiryCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModiryCartLogic) ModiryCart(in *pb.CartModiyRequest) (*pb.CartResponse, error) {
	key := fmt.Sprintf(l.svcCtx.Config.CachePreifx.Cart, in.MemberId)

	var list []*pb.Cart
	if err := redisx.GetAndJsonToObject(l.svcCtx.Redis, key, &list); err != nil {
		return nil, merrorx.NewCodeError(500, err.Error())
	}

	for _, value := range list {
		if value.Id == in.Id {
			value.Quantity = in.Quantity
			if l.svcCtx.CartModel.ModifyQuantity(value.Id, in.Quantity) == false {
				return nil, merrorx.NewCodeError(500, "修改失败")
			}

			break
		}
	}

	js, _ := json.Marshal(list)
	err := l.svcCtx.Redis.Set(key, string(js))

	if err != nil {
		return nil, merrorx.NewCodeError(500, "redis 错误")
	}

	return &pb.CartResponse{
		Code:   200,
		Detail: "ok",
	}, nil
}
