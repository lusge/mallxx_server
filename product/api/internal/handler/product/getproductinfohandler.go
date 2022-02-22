package product

import (
	"net/http"

	"mallxx_server/product/api/internal/logic/product"
	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetProductInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductInfoRequest
		if err := httpx.ParseForm(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := product.NewGetProductInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetProductInfo(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
