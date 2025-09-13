package middlewares

import (
	"fmt"
	"net/http"
	"strings"
)

type HPPOptions struct {
	CheckQuery                  bool
	CheckBody                   bool
	CheckBodyOnlyForContentType string
	Whitelist                   []string
}

func HPP(options HPPOptions) func(http.Handler) http.Handler {
	fmt.Println("HPP Middleware")
	return func(next http.Handler) http.Handler {
		fmt.Println("HPP Middleware being returned...")
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if options.CheckBody && r.Method == http.MethodPost && isCorrectcontent(r, options.CheckBodyOnlyForContentType) {
				// filter the body params
				FilterBodyParams(r, options.Whitelist)
			}
			if options.CheckQuery && r.URL.Query() != nil {
				// filter the query params
				FilterQueryParams(r, options.Whitelist)
			}
			next.ServeHTTP(w, r)
			fmt.Println("HPP Middleware ends...")
		})
	}
}

func isCorrectcontent(r *http.Request, contentType string) bool {
	return strings.Contains(r.Header.Get("Content-Type"), contentType)
}

func FilterBodyParams(r *http.Request, whitelist []string) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range r.Form {
		if len(v) > 1 {
			r.Form.Set(k, v[0])
			// r.Form.Set(k,v[len(v) - 1]) // last value
		}
		if !isWhiteListed(k, whitelist) {
			delete(r.Form, k)
		}
	}
}

func FilterQueryParams(r *http.Request, whitelist []string) {
	query := r.URL.Query()
	for k, v := range query {
		if len(v) > 1 {
			query.Set(k, v[0])
			// query.Form.Set(k,v[len(v) - 1]) // last value
		}
		if !isWhiteListed(k, whitelist) {
			query.Del(k)
		}
	}
	r.URL.RawQuery = query.Encode()
}

func isWhiteListed(params string, whitelist []string) bool {
	for _, v := range whitelist {
		if params == v {
			return true
		}
	}
	return false
}
