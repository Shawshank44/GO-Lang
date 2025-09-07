package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func ResponseTimeMiddleware(next http.Handler) http.Handler { // will calculate the time of response
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Sent Request in Response Time")
		start := time.Now()
		//Create a custom ResponseWriter to capture the status code
		wrappedWriter := &ResponseWriter{ResponseWriter: w, status: http.StatusOK}

		// Calculate the duration
		duration := time.Since(start)
		w.Header().Set("X-Response-Time", duration.String())
		next.ServeHTTP(wrappedWriter, r)
		// Log the request details :
		duration = time.Since(start)
		fmt.Printf("Method : %s, URL : %s, Status : %d, Duration : %v \n", r.Method, r.URL, wrappedWriter.status, duration.String())
		fmt.Println("Sent Response from Response Time Middleware")
	})
}

// Response writer :
type ResponseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *ResponseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}
