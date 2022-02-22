package category

import (
	"net/http"

	"mallxx_server/product/api/internal/logic/category"
	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCategoryInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdRequest
		if err := httpx.ParseForm(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := category.NewGetCategoryInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetCategoryInfo(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
