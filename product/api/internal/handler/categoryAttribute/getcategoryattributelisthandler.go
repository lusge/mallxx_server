package categoryAttribute

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mallxx_server/product/api/internal/logic/categoryAttribute"
	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
)

func GetCategoryAttributeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := categoryAttribute.NewGetCategoryAttributeListLogic(r.Context(), svcCtx)
		resp, err := l.GetCategoryAttributeList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
