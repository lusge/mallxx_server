package category

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mallxx_server/product/api/internal/logic/category"
	"mallxx_server/product/api/internal/svc"
)

func GetCategoryListWithChildrenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := category.NewGetCategoryListWithChildrenLogic(r.Context(), svcCtx)
		resp, err := l.GetCategoryListWithChildren()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
