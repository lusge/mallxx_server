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

type GetCartListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCartListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartListLogic {
	return &GetCartListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCartListLogic) GetCartList(in *pb.CartListRequest) (*pb.CartListResponse, error) {
	key := fmt.Sprintf(l.svcCtx.Config.CachePreifx.Cart, in.MemberId)
	list := make([]*pb.Cart, 0)

	if err := redisx.GetAndJsonToObject(l.svcCtx.Redis, key, &list); err != nil {
		return nil, merrorx.NewCodeError(500, err.Error())
	}

	if len(list) <= 0 {
		list = l.svcCtx.CartModel.FindAllByMemberId(in.MemberId)
		js, _ := json.Marshal(list)
		err := l.svcCtx.Redis.Set(key, string(js))
		fmt.Println(err, key, list)
	}

	return &pb.CartListResponse{
		Code:   200,
		Detail: "ok",
		Data:   list,
	}, nil
}
