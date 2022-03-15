package sign

import (
	"net/http"

	"mallxx_server/sign/api/internal/logic/sign"
	"mallxx_server/sign/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SignOutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := sign.NewSignOutLogic(r.Context(), svcCtx)
		resp, err := l.SignOut(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
