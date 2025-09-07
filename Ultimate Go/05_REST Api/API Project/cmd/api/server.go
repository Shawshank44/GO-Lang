package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	mw "restapi/internal/api/middlewares"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Welcome to the home page") // 1 way
	w.Write([]byte("Welcome to the home page"))
}

func TeachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Welcome to Teacher GET"))
		return
	case http.MethodPost:
		w.Write([]byte("Welcome to Teacher POST"))
		return
	case http.MethodPut:
		w.Write([]byte("Welcome to Teacher PUT"))
		return
	case http.MethodPatch:
		w.Write([]byte("Welcome to Teacher PATCH"))
		return
	case http.MethodDelete:
		w.Write([]byte("Welcome to Teacher DELETE"))
		return
	}

	w.Write([]byte("Welcome to Teacher Page"))
}

func StudentHandlers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Welcome to student GET"))
		return
	case http.MethodPost:
		w.Write([]byte("Welcome to student POST"))
		return
	case http.MethodPut:
		w.Write([]byte("Welcome to student PUT"))
		return
	case http.MethodPatch:
		w.Write([]byte("Welcome to student PATCH"))
		return
	case http.MethodDelete:
		w.Write([]byte("Welcome to student DELETE"))
		return
	}

	w.Write([]byte("Hello Welcome Student"))
}

func ExecutiveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Welcome to Executive GET"))
		return
	case http.MethodPost:
		w.Write([]byte("Welcome to Executive POST"))
		return
	case http.MethodPut:
		w.Write([]byte("Welcome to Executive PUT"))
		return
	case http.MethodPatch:
		w.Write([]byte("Welcome to Executive PATCH"))
		return
	case http.MethodDelete:
		w.Write([]byte("Welcome to Executive DELETE"))
		return
	}

	w.Write([]byte("Hello Welcome Executives"))
}

func main() {

	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux() //multiplexer (mux)

	mux.HandleFunc("/", RootHandler)
	mux.HandleFunc("/teachers/", TeachersHandler)
	mux.HandleFunc("/students/", StudentHandlers)
	mux.HandleFunc("/execs/", ExecutiveHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// Create custom server :
	server := &http.Server{
		Addr:      port,
		Handler:   mw.SecurityHeaders(mw.CORS(mux)),
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port ", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatal(err)
	}
}
