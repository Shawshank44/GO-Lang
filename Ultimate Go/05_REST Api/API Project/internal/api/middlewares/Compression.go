package middlewares

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
)

func Compression(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the client accepts Gzip encoding
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
		}
		// Set the response Header
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		// Write the ResponseWriter
		w = &gzipResponseWriter{ResponseWriter: w, Writer: gz}

		next.ServeHTTP(w, r)
		fmt.Println("Sent Response From Copression Middleware")
	})
}

// GzipResponseWriter wraps http.ResponseWriter to write Gzipped Responses
type gzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (g *gzipResponseWriter) Write(b []byte) (int, error) {
	return g.Writer.Write(b)
}
