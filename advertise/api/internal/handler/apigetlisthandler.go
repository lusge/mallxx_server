package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mallxx_server/advertise/api/internal/logic"
	"mallxx_server/advertise/api/internal/svc"
	"mallxx_server/advertise/api/internal/types"
)

func apiGetListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewApiGetListLogic(r.Context(), svcCtx)
		resp, err := l.ApiGetList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
