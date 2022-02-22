package recommend

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mallxx_server/product/api/internal/logic/recommend"
	"mallxx_server/product/api/internal/svc"
	"mallxx_server/product/api/internal/types"
)

func DeleteRecommendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecommendRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := recommend.NewDeleteRecommendLogic(r.Context(), svcCtx)
		resp, err := l.DeleteRecommend(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
