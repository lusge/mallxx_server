package interceptor

import (
	"context"
	"mallxx_server/common/merrorx"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
)

//改了也没用，啥玩意
func ErrorInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	resp, err = handler(ctx, req)
	if err != nil {
		logx.WithContext(ctx).Errorf("【PRC-ERROR】 %+v", err)
		if e, ok := err.(*merrorx.CodeError); ok {
			return http.StatusOK, e
		} else {
			ret, ok := merrorx.ConvertRpcError(err)
			if ok {
				return http.StatusOK, ret
			}

		}
		return http.StatusOK, &merrorx.CodeError{
			Code:   500,
			Detail: err.Error(),
		}
	}

	return resp, err
}
