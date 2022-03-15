package cart

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mallxx_server/cart/api/internal/logic/cart"
	"mallxx_server/cart/api/internal/svc"
	"mallxx_server/cart/api/internal/types"
)

func ModifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CartModiyRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := cart.NewModifyLogic(r.Context(), svcCtx)
		resp, err := l.Modify(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
