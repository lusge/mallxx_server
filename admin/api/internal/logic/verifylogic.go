package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"mallxx_server/admin/api/internal/svc"
	"mallxx_server/admin/rpc/adminservice"
	"mallxx_server/common/merrorx"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	ztoken "github.com/zeromicro/go-zero/rest/token"
)

type VerifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) VerifyLogic {
	return VerifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyLogic) Verify(r *http.Request) (resp int64, err error) {
	// todo: add your logic here and delete this line
	// token := r.Header.Get("Authorization")
	requestPath := r.Header.Get("X-Original-Uri")
	if strings.Contains(requestPath, "?") {
		requestPath = strings.Split(requestPath, "?")[0]
	}

	if requestPath == "/admin/login" {
		return 0, nil
	}

	parser := ztoken.NewTokenParser()
	tok, err := parser.ParseToken(r, l.svcCtx.Config.JwtAuth.AccessSecret, "")

	if err != nil {
		return 0, merrorx.NewCodeError(500, "token verify failed")
	}
	if tok.Valid {
		claims, ok := tok.Claims.(jwt.MapClaims) // 解析token中对内容
		if ok {
			fmt.Print(claims)
			userId, _ := claims["uid"].(json.Number).Int64()

			if userId <= 0 {
				return 0, merrorx.NewCodeError(500, "token verify failed")
			}

			verifyResp, err := l.svcCtx.AdminRpc.VerifyToken(l.ctx, &adminservice.AdminTokenRequest{
				Token: tok.Raw,
				Id:    int64(userId),
			})

			if err != nil {
				return 0, err
			}

			if verifyResp.Code == 200 {
				return int64(userId), nil
			}
		}
	}

	return 0, merrorx.NewCodeError(500, "token verify failed")
}
