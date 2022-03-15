package cart

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mallxx_server/cart/api/internal/logic/cart"
	"mallxx_server/cart/api/internal/svc"
	"mallxx_server/cart/api/internal/types"
)

func AddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CartAddRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := cart.NewAddLogic(r.Context(), svcCtx)
		resp, err := l.Add(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
