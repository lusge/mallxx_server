package sku

import (
	"net/http"

	"mallxx_server/product/api/internal/logic/sku"
	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetSkuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SkuStockRequest
		if err := httpx.ParseForm(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := sku.NewGetSkuLogic(r.Context(), svcCtx)
		resp, err := l.GetSku(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
