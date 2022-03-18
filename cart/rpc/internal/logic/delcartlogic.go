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

type DelCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCartLogic {
	return &DelCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCartLogic) DelCart(in *pb.CartDelRequest) (*pb.CartResponse, error) {
	key := fmt.Sprintf(l.svcCtx.Config.CachePreifx.Cart, in.MemberId)
	var list []*pb.Cart
	if err := redisx.GetAndJsonToObject(l.svcCtx.Redis, key, &list); err != nil {
		fmt.Println(err, "kkk", key)
		return nil, merrorx.NewCodeError(500, err.Error())
	}

	if len(in.Ids) <= 0 {
		return nil, merrorx.NewCodeError(500, "没有选择")
	}

	for index, value := range list {
		for _, id := range in.Ids {
			if value.Id == id {
				if len(list) == 1 {
					list = nil
				} else {
					list = append(list[:index], list[index+1:]...)
				}

				if !l.svcCtx.CartModel.Delete(value.Id) {
					return nil, merrorx.NewCodeError(500, "删除是吧")
				} else {
					if len(list) == 0 {
						l.svcCtx.Redis.Del(key)
					} else {
						js, _ := json.Marshal(list)
						_ = l.svcCtx.Redis.Set(key, string(js))
					}

					return &pb.CartResponse{
						Code:   200,
						Detail: "ok",
					}, nil
				}
			}
		}

	}
	return nil, merrorx.NewCodeError(500, "删除失败")
}
