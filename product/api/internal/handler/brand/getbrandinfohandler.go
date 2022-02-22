package brand

import (
	"net/http"

	"mallxx_server/product/api/internal/logic/brand"
	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetBrandInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BrandRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := brand.NewGetBrandInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetBrandInfo(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
