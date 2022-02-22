package sku

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mallxx_server/product/api/internal/logic/sku"
	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
)

func UpdateSkuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SkuStockListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := sku.NewUpdateSkuLogic(r.Context(), svcCtx)
		resp, err := l.UpdateSku(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
