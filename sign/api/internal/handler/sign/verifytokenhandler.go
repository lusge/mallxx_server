package sign

import (
	"net/http"
	"strconv"

	"mallxx_server/sign/api/internal/logic/sign"
	"mallxx_server/sign/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func VerifyTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := sign.NewVerifyTokenLogic(r.Context(), svcCtx)
		resp, err := l.VerifyToken(r)

		w.Header().Set("x-user", strconv.FormatInt(resp, 10))

		if err != nil {
			httpx.WriteJson(w, http.StatusUnauthorized, err.Error())
		} else {
			httpx.WriteJson(w, http.StatusOK, "ok")
		}

	}
}
