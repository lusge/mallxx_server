package handler

import (
	"net/http"
	"strconv"

	"mallxx_server/admin/api/internal/logic"
	"mallxx_server/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func verifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewVerifyLogic(r.Context(), svcCtx)
		resp, err := l.Verify(r)

		w.Header().Set("x-user", strconv.FormatInt(resp, 10))

		if err != nil {
			httpx.WriteJson(w, http.StatusUnauthorized, err.Error())
		}

		httpx.WriteJson(w, http.StatusOK, "ok")
	}
}
