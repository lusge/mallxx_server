package metadata

import "context"

const UIDKEY = "uid"

func GetUidByCtx(ctx context.Context) int64 {
	uid, _ := ctx.Value(UIDKEY).(int64)
	return uid
}
