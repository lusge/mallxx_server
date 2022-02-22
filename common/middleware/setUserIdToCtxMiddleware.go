package middleware

import (
	"context"
	"net/http"
	"strconv"
)

func SetUserId(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, _ := strconv.ParseInt(r.Header.Get("X-User"), 10, 64)
		ctx := r.Context()
		ctx = context.WithValue(ctx, "uid", userId)

		next(w, r.WithContext(ctx))
	}
}
