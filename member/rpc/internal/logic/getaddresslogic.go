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

type GetAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAddressLogic {
	return &GetAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAddressLogic) GetAddress(in *pb.MemberRequest) (*pb.ReceiveAddressListResponse, error) {
	key := fmt.Sprintf(l.svcCtx.Config.CachePreifx.Address, in.Uid)
	var list []*pb.ReceiveAddress
	if err := redisx.GetAndJsonToObject(l.svcCtx.Redis, key, &list); err != nil {
		return nil, merrorx.NewCodeError(500, err.Error())
	}

	if len(list) <= 0 {
		list = l.svcCtx.ReceiveAddressModel.FindAllByMemberId(in.Uid)
		if !l.isThereDefault(list) && len(list) > 0 {
			list[0].DefaultStatus = 1
		}

		js, _ := json.Marshal(list)
		err := l.svcCtx.Redis.Set(key, string(js))

		if err != nil {
			return nil, merrorx.NewCodeError(500, "redis 错误")
		}
	}

	return &pb.ReceiveAddressListResponse{
		Data:   list,
		Code:   200,
		Detail: "ok",
	}, nil
}

func (l *GetAddressLogic) isThereDefault(list []*pb.ReceiveAddress) bool {
	for _, value := range list {
		if value.DefaultStatus == 1 {
			return true
		}
	}
	return false
}
