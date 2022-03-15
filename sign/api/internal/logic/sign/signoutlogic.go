package sign

import (
	"context"
	"encoding/json"
	"net/http"

	"mallxx_server/common/merrorx"
	"mallxx_server/sign/api/internal/svc"
	"mallxx_server/sign/api/internal/types"
	"mallxx_server/sign/rpc/signservice"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	ztoken "github.com/zeromicro/go-zero/rest/token"
)

type SignOutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) SignOutLogic {
	return SignOutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignOutLogic) SignOut(r *http.Request) (*types.Response, error) {
	parser := ztoken.NewTokenParser()
	tok, err := parser.ParseToken(r, l.svcCtx.Config.JwtAuth.AccessSecret, "")

	if err != nil {
		return nil, merrorx.NewCodeError(500, err.Error())
	}

	if tok.Valid {
		claims, ok := tok.Claims.(jwt.MapClaims) // 解析token中对内容

		if ok {
			userId, err := claims["uid"].(json.Number).Int64()

			if err != nil {
				return nil, merrorx.NewCodeError(500, err.Error())
			}

			if userId <= 0 {
				return nil, merrorx.NewCodeError(500, "not found user id")
			}

			verifyResp, err := l.svcCtx.SignRpc.CleanToken(l.ctx, &signservice.SignTokenRequest{
				Token: tok.Raw,
				Id:    int64(userId),
			})

			if err != nil {
				return nil, err
			}

			if verifyResp.Code == 200 {
				return &types.Response{
					Code:   200,
					Detail: "ok",
				}, nil
			}
		} else {
			return nil, merrorx.NewCodeError(500, "token Claims failed")
		}
	}

	return nil, merrorx.NewCodeError(500, "token verify failed")
}
