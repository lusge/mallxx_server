package address

import (
	"net/http"

	"mallxx_server/member/api/internal/logic/address"
	"mallxx_server/member/api/internal/svc"
	"mallxx_server/member/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetAddressInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReceiveAddressRequest
		if err := httpx.ParseForm(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := address.NewGetAddressInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetAddressInfo(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
