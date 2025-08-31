package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprint(w, "Welcome to the home page") // 1 way
		w.Write([]byte("Welcome to the home page"))
	})
	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
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
	})
	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
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
	})
	http.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {
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
	})

	port := ":3000"
	fmt.Println("Server is running on port ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
