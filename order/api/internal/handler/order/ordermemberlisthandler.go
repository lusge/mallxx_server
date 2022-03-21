package order

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mallxx_server/order/api/internal/logic/order"
	"mallxx_server/order/api/internal/svc"
	"mallxx_server/order/api/internal/types"
)

func OrderMemberListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderMemberRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := order.NewOrderMemberListLogic(r.Context(), svcCtx)
		resp, err := l.OrderMemberList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
