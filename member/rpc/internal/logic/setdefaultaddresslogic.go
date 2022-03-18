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

type SetDefaultAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetDefaultAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetDefaultAddressLogic {
	return &SetDefaultAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetDefaultAddressLogic) SetDefaultAddress(in *pb.ReceiveAddressRequest) (*pb.MemberResponse, error) {
	key := fmt.Sprintf(l.svcCtx.Config.CachePreifx.Address, in.MemberId)
	var list []*pb.ReceiveAddress
	if err := redisx.GetAndJsonToObject(l.svcCtx.Redis, key, &list); err != nil {
		return nil, merrorx.NewCodeError(500, err.Error())
	}

	if !l.svcCtx.ReceiveAddressModel.SetDefault(in.MemberId, in.Id) {
		return nil, merrorx.NewCodeError(500, "设置失败")
	}

	for index, value := range list {
		if value.Id == in.Id {
			list[index].DefaultStatus = 1
			js, _ := json.Marshal(list)
			err := l.svcCtx.Redis.Set(key, string(js))

			if err != nil {
				return nil, merrorx.NewCodeError(500, "redis 错误")
			}

			return &pb.MemberResponse{
				Code:   200,
				Detail: "设置成功",
			}, nil
		}
	}
	return nil, merrorx.NewCodeError(500, "设置失败")
}
