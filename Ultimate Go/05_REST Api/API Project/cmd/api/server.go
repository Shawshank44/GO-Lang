package main

import (
	"fmt"
	"log"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Welcome to the home page") // 1 way
	w.Write([]byte("Welcome to the home page"))
}

func TeachersHandler(w http.ResponseWriter, r *http.Request) {
	// teachers/{id} -> PathParams
	// teachers/?key=value&query=value2&sortby=email&sortorder=ASC -> QueryParams
	switch r.Method {
	case http.MethodGet:
		// PathParams :
		// fmt.Println(r.URL.Path)
		// path := strings.TrimPrefix(r.URL.Path, "/teachers/")
		// userID := strings.TrimSuffix(path, "/")
		// fmt.Println("The ID is : ", userID)

		// QueryParams :
		fmt.Println("Query Params", r.URL.Query())
		queryParams := r.URL.Query()
		sortby := queryParams.Get("sortby")
		key := queryParams.Get("key")
		sortOrder := queryParams.Get("sortorder")

		if sortOrder == "" { //If Blank
			sortOrder = "DESC"
		}
		fmt.Printf("Sortby : %v, SortOrder :%v, key : %v \n", sortby, sortOrder, key)

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
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/teachers/", TeachersHandler)
	http.HandleFunc("/students/", StudentHandlers)
	http.HandleFunc("/execs/", ExecutiveHandler)

	port := ":3000"
	fmt.Println("Server is running on port ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
