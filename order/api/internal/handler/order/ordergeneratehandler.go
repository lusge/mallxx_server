package order

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mallxx_server/order/api/internal/logic/order"
	"mallxx_server/order/api/internal/svc"
	"mallxx_server/order/api/internal/types"
)

func OrderGenerateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GenerateOrderRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := order.NewOrderGenerateLogic(r.Context(), svcCtx)
		resp, err := l.OrderGenerate(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
