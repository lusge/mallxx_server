package handler

import (
	"net/http"

	"mallxx_server/advertise/api/internal/logic"
	"mallxx_server/advertise/api/internal/svc"
	"mallxx_server/advertise/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func infoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdOneRequest
		if err := httpx.ParseForm(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewInfoLogic(r.Context(), svcCtx)
		resp, err := l.Info(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
