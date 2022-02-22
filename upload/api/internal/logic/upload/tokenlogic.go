package upload

import (
	"context"

	"mallxx_server/common/qiniu"
	"mallxx_server/upload/api/internal/svc"
	"mallxx_server/upload/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) TokenLogic {
	return TokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TokenLogic) Token() (resp *types.Response, err error) {
	token := qiniu.UploadToken(l.svcCtx.Config.Qiniu.AccessKey, l.svcCtx.Config.Qiniu.SecretKey, l.svcCtx.Config.Qiniu.Bucket)
	return &types.Response{
		Code:   200,
		Token:  token,
		Detail: "ok",
		Host:   l.svcCtx.Config.Qiniu.Host,
	}, nil
}
