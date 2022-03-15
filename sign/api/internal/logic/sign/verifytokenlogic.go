package sign

import (
	"context"
	"encoding/json"
	"net/http"

	"mallxx_server/common/merrorx"
	"mallxx_server/sign/api/internal/svc"
	"mallxx_server/sign/rpc/signservice"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	ztoken "github.com/zeromicro/go-zero/rest/token"
)

type VerifyTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) VerifyTokenLogic {
	return VerifyTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyTokenLogic) VerifyToken(r *http.Request) (int64, error) {
	parser := ztoken.NewTokenParser()
	tok, err := parser.ParseToken(r, l.svcCtx.Config.JwtAuth.AccessSecret, "")

	if err != nil {
		return 0, merrorx.NewCodeError(500, err.Error())
	}

	if tok.Valid {
		claims, ok := tok.Claims.(jwt.MapClaims) // 解析token中对内容

		if ok {
			userId, err := claims["uid"].(json.Number).Int64()

			if err != nil {
				return 0, merrorx.NewCodeError(500, err.Error())
			}

			if userId <= 0 {
				return 0, merrorx.NewCodeError(500, "not found user id")
			}

			verifyResp, err := l.svcCtx.SignRpc.VerifyToken(l.ctx, &signservice.SignTokenRequest{
				Token: tok.Raw,
				Id:    int64(userId),
			})

			if err != nil {
				return 0, err
			}

			if verifyResp.Code == 200 {
				return int64(userId), nil
			}
		} else {
			return 0, merrorx.NewCodeError(500, "token Claims failed")
		}
	}

	return 0, merrorx.NewCodeError(500, "token verify failed")
}
