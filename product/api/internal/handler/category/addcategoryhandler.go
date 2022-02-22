package category

import (
	"fmt"
	"net/http"

	"mallxx_server/product/api/internal/logic/category"
	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Category
		if err := httpx.Parse(r, &req); err != nil {
			fmt.Println(err)
			httpx.Error(w, err)
			return
		}

		l := category.NewAddCategoryLogic(r.Context(), svcCtx)
		resp, err := l.AddCategory(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
