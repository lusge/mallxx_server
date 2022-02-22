package recommend

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mallxx_server/product/api/internal/logic/recommend"
	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
)

func SetRecommendSortHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecommendSetSortRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := recommend.NewSetRecommendSortLogic(r.Context(), svcCtx)
		resp, err := l.SetRecommendSort(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
