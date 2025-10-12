package handlers

import "net/http"

func RootHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Welcome to the home page") // 1 way
	w.Write([]byte("Welcome to the School"))
}
