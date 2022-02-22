package productAttribute

import (
	"net/http"

	"mallxx_server/product/api/internal/logic/productAttribute"
	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteProductAttributeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdRequest
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := productAttribute.NewDeleteProductAttributeLogic(r.Context(), svcCtx)
		resp, err := l.DeleteProductAttribute(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
