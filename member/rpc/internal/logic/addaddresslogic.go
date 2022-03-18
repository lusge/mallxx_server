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

type AddAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAddressLogic {
	return &AddAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddAddressLogic) AddAddress(in *pb.ReceiveAddress) (*pb.ReceiveAddressResponse, error) {
	key := fmt.Sprintf(l.svcCtx.Config.CachePreifx.Address, in.MemberId)
	var list []*pb.ReceiveAddress
	if err := redisx.GetAndJsonToObject(l.svcCtx.Redis, key, &list); err != nil {
		return nil, merrorx.NewCodeError(500, err.Error())
	}

	if l.svcCtx.ReceiveAddressModel.Insert(in) == false {
		return nil, merrorx.NewCodeError(500, "新增失败")
	}

	list = append(list, in)
	js, _ := json.Marshal(list)
	err := l.svcCtx.Redis.Set(key, string(js))

	if err != nil {
		return nil, merrorx.NewCodeError(500, "redis 错误")
	}

	return &pb.ReceiveAddressResponse{
		Code:   200,
		Detail: "ok",
		Data:   in,
	}, nil
}
