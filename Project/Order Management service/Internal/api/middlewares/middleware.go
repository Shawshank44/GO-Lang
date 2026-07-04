package middlewares

import "net/http"

type MiddleWare func(http.Handler) http.Handler

func ApplyMiddleWares(handler http.Handler, middleWares ...MiddleWare) http.Handler {
	for _, middleware := range middleWares {
		handler = middleware(handler)
	}

	return handler
}
