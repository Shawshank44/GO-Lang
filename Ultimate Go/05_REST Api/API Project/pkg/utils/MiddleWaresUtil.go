package utils

import "net/http"

// Middleware is a function that wraps an http.Handler with additional functionality
type MiddleWare func(http.Handler) http.Handler

func ApplyMiddleWares(handler http.Handler, middlewares ...MiddleWare) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
