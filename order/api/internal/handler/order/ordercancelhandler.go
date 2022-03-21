package order

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mallxx_server/order/api/internal/logic/order"
	"mallxx_server/order/api/internal/svc"
	"mallxx_server/order/api/internal/types"
)

func OrderCancelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := order.NewOrderCancelLogic(r.Context(), svcCtx)
		resp, err := l.OrderCancel(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
