package handler

import (
	"fmt"
	"net/http"

	"mallxx_server/advertise/api/internal/logic"
	"mallxx_server/advertise/api/internal/svc"
	"mallxx_server/advertise/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func adminGetListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdRequest
		if err := httpx.ParseForm(r, &req); err != nil {
			fmt.Println(err)
			httpx.Error(w, err)
			return
		}

		l := logic.NewAdminGetListLogic(r.Context(), svcCtx)
		resp, err := l.AdminGetList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
