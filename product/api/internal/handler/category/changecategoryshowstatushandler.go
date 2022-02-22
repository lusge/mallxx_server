package category

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mallxx_server/product/api/internal/logic/category"
	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
)

func ChangeCategoryShowStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CategoryChangeStatus
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := category.NewChangeCategoryShowStatusLogic(r.Context(), svcCtx)
		resp, err := l.ChangeCategoryShowStatus(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
