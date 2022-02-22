package categoryAttribute

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mallxx_server/product/api/internal/logic/categoryAttribute"
	"mallxx_server/product/api/internal/svc"
)

func GetCategoryAttributeAndProductAttributeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := categoryAttribute.NewGetCategoryAttributeAndProductAttributeLogic(r.Context(), svcCtx)
		resp, err := l.GetCategoryAttributeAndProductAttribute()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
