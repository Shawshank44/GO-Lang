package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

func LoadClientCAs() *x509.CertPool {
	clientCAs := x509.NewCertPool()
	caCert, err := os.ReadFile("cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	clientCAs.AppendCertsFromPEM(caCert)
	return clientCAs
}

func main() {

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		LogRequestDetails(r)
		fmt.Fprintf(w, "Handeling incoming orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		LogRequestDetails(r)
		fmt.Fprintf(w, "Handeling users")
	})

	port := 3000

	// Load the TLS cert and key
	cert := "cert.pem"
	key := "key.pem"

	// Configure TLS
	tlsconfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		ClientAuth: tls.RequireAndVerifyClientCert, // Enforce mTLS
		ClientCAs:  LoadClientCAs(),
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

func LogRequestDetails(r *http.Request) {
	httpVersion := r.Proto
	fmt.Println("Received request with http version : ", httpVersion)
	if r.TLS != nil {
		tlsVersion := GetTLSVersion(r.TLS.Version)
		fmt.Println("Received request with TLS version", tlsVersion)
	} else {
		fmt.Println("Received request without TLS")
	}
}

func GetTLSVersion(version uint16) string {
	switch version {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return "Unknown TLS version"
	}
}
