package middlewares

import (
	"net/http"
	"strings"
)

func MiddlewaresExcludeParts(middleware func(http.Handler) http.Handler, excludePaths ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, Path := range excludePaths {
				if strings.HasPrefix(r.URL.Path, Path) {
					next.ServeHTTP(w, r)
					return
				}
			}
			middleware(next).ServeHTTP(w, r)
		})
	}
}
