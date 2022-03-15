package cart

import (
	"net/http"

	"mallxx_server/cart/api/internal/logic/cart"
	"mallxx_server/cart/api/internal/svc"
	"mallxx_server/cart/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CartListRequest
		if err := httpx.ParseForm(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := cart.NewGetListLogic(r.Context(), svcCtx)
		resp, err := l.GetList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
