package handler

import (
	"fmt"
	"net/http"

	"mallxx_server/advertise/api/internal/logic"
	"mallxx_server/advertise/api/internal/svc"
	"mallxx_server/advertise/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func deleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdOneRequest
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			fmt.Println(err, "err")
			httpx.Error(w, err)
			return
		}

		l := logic.NewDeleteLogic(r.Context(), svcCtx)
		resp, err := l.Delete(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
