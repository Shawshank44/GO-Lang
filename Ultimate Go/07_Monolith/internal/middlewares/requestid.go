package middlewares

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type ctxKey int

const (
	requestIDKey ctxKey = iota
)

const (
	requestID = "X-Request-ID"
)

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get(requestID)
		if id == "" {
			id = uuid.NewString()
		}
		w.Header().Add(requestID, id)

		ctx := context.WithValue(r.Context(), requestIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RequestIDFromContext(ctx context.Context) string {
	requestID := ctx.Value(requestIDKey).(string)
	return requestID
}
