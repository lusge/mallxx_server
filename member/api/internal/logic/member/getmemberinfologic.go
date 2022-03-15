package member

import (
	"context"

	"mallxx_server/common/merrorx"
	"mallxx_server/member/api/internal/svc"
	"mallxx_server/member/api/internal/types"
	"mallxx_server/member/rpc/memberservice"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMemberInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetMemberInfoLogic {
	return GetMemberInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMemberInfoLogic) GetMemberInfo(req types.MemberRequest) (*types.MemberResponse, error) {
	userId := l.ctx.Value("uid").(int64)
	if userId <= 0 {
		return nil, merrorx.NewCodeError(500, "参数错误")
	}

	resp, err := l.svcCtx.MemberRpc.GetMemberInfo(l.ctx, &memberservice.MemberRequest{
		Uid: userId,
	})

	if err != nil {
		return nil, err
	}

	var data types.MemberData
	copier.Copy(&data, resp.Data)
	return &types.MemberResponse{
		Code:   resp.Code,
		Detail: resp.Detail,
		Data:   &data,
	}, nil
}
