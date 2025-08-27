package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handeling incoming orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handeling users")
	})

	port := 3000

	// Load the TLS cert and key
	cert := "cert.pem"
	key := "key.pem"

	// Configure TLS
	tlsconfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// Create a custom server
	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		Handler:   nil,
		TLSConfig: tlsconfig,
	}

	// Enable http2
	http2.ConfigureServer(server, &http2.Server{})

	fmt.Printf("Server running in : %d \n", port)

	// Http2 server (use https connection) :
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatal(err)
	}

	// Http server (use http connection):
	// fmt.Printf("Server running in http://127.0.0.1:%d \n", port)
	// err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
