package address

import (
	"net/http"

	"mallxx_server/member/api/internal/logic/address"
	"mallxx_server/member/api/internal/svc"
	"mallxx_server/member/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetAddressListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MemberRequest
		userId := r.Context().Value("uid").(int64)
		req.Uid = userId

		l := address.NewGetAddressListLogic(r.Context(), svcCtx)
		resp, err := l.GetAddressList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
