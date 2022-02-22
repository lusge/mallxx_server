package brand

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mallxx_server/product/api/internal/logic/brand"
	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
)

func UpdateBrandShowStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BrandStatusRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := brand.NewUpdateBrandShowStatusLogic(r.Context(), svcCtx)
		resp, err := l.UpdateBrandShowStatus(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
