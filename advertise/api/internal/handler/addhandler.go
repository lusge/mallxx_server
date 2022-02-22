package handler

import (
	"fmt"
	"net/http"

	"mallxx_server/advertise/api/internal/logic"
	"mallxx_server/advertise/api/internal/svc"
	"mallxx_server/advertise/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func addHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Advertise
		if err := httpx.Parse(r, &req); err != nil {
			fmt.Println(err)
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddLogic(r.Context(), svcCtx)
		resp, err := l.Add(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
