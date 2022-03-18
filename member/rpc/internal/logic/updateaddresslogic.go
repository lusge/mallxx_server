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

type UpdateAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAddressLogic {
	return &UpdateAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAddressLogic) UpdateAddress(in *pb.ReceiveAddress) (*pb.MemberResponse, error) {
	key := fmt.Sprintf(l.svcCtx.Config.CachePreifx.Address, in.MemberId)
	var list []*pb.ReceiveAddress
	if err := redisx.GetAndJsonToObject(l.svcCtx.Redis, key, &list); err != nil {
		return nil, merrorx.NewCodeError(500, err.Error())
	}

	if !l.svcCtx.ReceiveAddressModel.Update(in) {
		return nil, merrorx.NewCodeError(500, "更新失败")
	}

	for index, value := range list {
		if value.Id == in.Id {
			list[index] = in
		}
	}

	js, _ := json.Marshal(list)
	err := l.svcCtx.Redis.Set(key, string(js))

	if err != nil {
		return nil, merrorx.NewCodeError(500, "redis 错误")
	}

	return &pb.MemberResponse{
		Code:   200,
		Detail: "更新成功",
	}, nil
}
