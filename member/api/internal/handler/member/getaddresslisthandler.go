package member

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mallxx_server/member/api/internal/logic/member"
	"mallxx_server/member/api/internal/svc"
	"mallxx_server/member/api/internal/types"
)

func GetAddressListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MemberRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := member.NewGetAddressListLogic(r.Context(), svcCtx)
		resp, err := l.GetAddressList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}