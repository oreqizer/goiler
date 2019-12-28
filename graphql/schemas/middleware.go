package schemas

import (
	"context"
	"net/http"
)

type key struct {
	name string
}

var keyAccount = &key{"account"}

// Middleware creates the loader middleware
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			ctx = context.WithValue(ctx, keyAccount, MakeAccountLoader(ctx))

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
