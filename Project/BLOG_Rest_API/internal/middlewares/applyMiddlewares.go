package middlewares

import "net/http"

type MiddleWare func(http.Handler) http.Handler

func ApplyMiddleWares(handler http.Handler, middlewares ...MiddleWare) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
