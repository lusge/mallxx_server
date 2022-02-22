package category

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mallxx_server/product/api/internal/logic/category"
	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
)

func UpdateCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Category
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := category.NewUpdateCategoryLogic(r.Context(), svcCtx)
		resp, err := l.UpdateCategory(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
