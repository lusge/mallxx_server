package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"mallxx_server/common/merrorx"
	"mallxx_server/common/redisx"
	"mallxx_server/member/rpc/internal/svc"
	"mallxx_server/member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAddressLogic {
	return &DelAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAddressLogic) DelAddress(in *pb.ReceiveAddressRequest) (*pb.MemberResponse, error) {
	key := fmt.Sprintf(l.svcCtx.Config.CachePreifx.Address, in.MemberId)
	var list []*pb.ReceiveAddress
	if err := redisx.GetAndJsonToObject(l.svcCtx.Redis, key, &list); err != nil {
		return nil, merrorx.NewCodeError(500, err.Error())
	}

	for index, value := range list {
		if value.Id == in.Id {
			if len(list) == 1 {
				list = nil
			} else {
				list = append(list[:index], list[index+1:]...)
			}

			if !l.svcCtx.ReceiveAddressModel.Delete(in.MemberId, in.Id) {
				return nil, merrorx.NewCodeError(500, "删除失败")
			} else {
				if len(list) == 0 {
					l.svcCtx.Redis.Del(key)
				} else {
					js, _ := json.Marshal(list)
					_ = l.svcCtx.Redis.Set(key, string(js))
				}

				return &pb.MemberResponse{
					Code:   200,
					Detail: "删除成功",
				}, nil
			}

		}
	}

	return nil, merrorx.NewCodeError(500, "删除失败")
}
