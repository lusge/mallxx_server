package logic

import (
	"context"
	"fmt"

	"mallxx_server/common/merrorx"
	"mallxx_server/common/redisx"
	"mallxx_server/member/rpc/internal/svc"
	"mallxx_server/member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAddressInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAddressInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAddressInfoLogic {
	return &GetAddressInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAddressInfoLogic) GetAddressInfo(in *pb.ReceiveAddressRequest) (*pb.ReceiveAddressResponse, error) {
	key := fmt.Sprintf(l.svcCtx.Config.CachePreifx.Address, in.MemberId)
	var list []*pb.ReceiveAddress
	if err := redisx.GetAndJsonToObject(l.svcCtx.Redis, key, &list); err != nil {
		return nil, merrorx.NewCodeError(500, err.Error())
	}

	var address *pb.ReceiveAddress

	if len(list) <= 0 {
		if in.Id <= 0 {
			address = l.svcCtx.ReceiveAddressModel.FindOneByMemberIdWithDefault(in.MemberId)
		} else {
			address = l.svcCtx.ReceiveAddressModel.FindOneById(in.Id)
		}
	} else {
		for index, value := range list {
			if index == 0 {
				address = value
			}

			if in.Id <= 0 {
				if value.DefaultStatus == 1 {
					address = value
				}
			} else if value.Id == in.Id {
				address = value
			}
		}
	}

	return &pb.ReceiveAddressResponse{
		Code:   200,
		Detail: "ok",
		Data:   address,
	}, nil
}
