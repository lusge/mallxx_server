package categoryAttribute

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mallxx_server/product/api/internal/logic/categoryAttribute"
	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
)

func DelCategoryAttributeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CategoryAttribute
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := categoryAttribute.NewDelCategoryAttributeLogic(r.Context(), svcCtx)
		resp, err := l.DelCategoryAttribute(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
