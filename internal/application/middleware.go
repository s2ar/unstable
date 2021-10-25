package application

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

func WithApp(app Application) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), ContextApp, app)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
