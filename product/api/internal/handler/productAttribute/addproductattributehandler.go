package productAttribute

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mallxx_server/product/api/internal/logic/productAttribute"
	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
)

func AddProductAttributeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductAttribute
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := productAttribute.NewAddProductAttributeLogic(r.Context(), svcCtx)
		resp, err := l.AddProductAttribute(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
