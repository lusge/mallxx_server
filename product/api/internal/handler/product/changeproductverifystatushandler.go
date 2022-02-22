package product

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mallxx_server/product/api/internal/logic/product"
	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
)

func ChangeProductVerifyStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductChangeStatusRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := product.NewChangeProductVerifyStatusLogic(r.Context(), svcCtx)
		resp, err := l.ChangeProductVerifyStatus(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
