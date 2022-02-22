package interceptor

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
)

//改了也没用，啥玩意
func ErrorInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	resp, err = handler(ctx, req)
	if err != nil {
		logx.WithContext(ctx).Errorf("【PRC-ERROR】 %+v", err)
		//转成自定义 err
		// err = merrorx.ConvertRpcError(err)
	}

	return resp, err
}
