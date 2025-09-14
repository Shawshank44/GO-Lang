package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	mw "restapi/internal/api/middlewares"
	"strconv"
	"strings"
	"sync"
)

type Teacher struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Class     string `json:"class"`
	Subject   string `json:"subject"`
}

var (
	teachers = make(map[int]Teacher)
	mutex    = &sync.Mutex{}
	nextId   = 1
)

func init() {
	teachers[nextId] = Teacher{
		ID:        nextId,
		FirstName: "John",
		LastName:  "Doe",
		Class:     "9A",
		Subject:   "Math",
	}
	nextId++
	teachers[nextId] = Teacher{
		ID:        nextId,
		FirstName: "Jane",
		LastName:  "smith",
		Class:     "10A",
		Subject:   "Algebra",
	}
	nextId++
	teachers[nextId] = Teacher{
		ID:        nextId,
		FirstName: "Jane",
		LastName:  "Doe",
		Class:     "10B",
		Subject:   "Drawing",
	}
	nextId++
	teachers[nextId] = Teacher{
		ID:        nextId,
		FirstName: "Antheny",
		LastName:  "Missery",
		Class:     "10C",
		Subject:   "Science",
	}
	nextId++
}

func GetTeachersHandeler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idstr := strings.TrimSuffix(path, "/")
	fmt.Println(idstr)

	if idstr == "" {
		firstName := r.URL.Query().Get("first_name") // encoding/json package automatically converts your exported field names into JSON keys
		lastName := r.URL.Query().Get("last_name")

		TeacherList := make([]Teacher, 0, len(teachers))
		for _, teacher := range teachers {
			if (firstName == "" || teacher.FirstName == firstName) && (lastName == "" || teacher.LastName == lastName) {
				TeacherList = append(TeacherList, teacher)
			}
		}
		response := struct {
			Status string    `json:"status"`
			Count  int       `json:"count"`
			Data   []Teacher `json:"data"`
		}{
			Status: "Success",
			Count:  len(teachers),
			Data:   TeacherList,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

	// Handle Path parameter :
	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println(err)
		return
	}
	teacher, exists := teachers[id]
	if !exists {
		http.Error(w, "Teacher not found", http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(teacher)
}

func AddTeacherHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	var NewTeachers []Teacher
	err := json.NewDecoder(r.Body).Decode(&NewTeachers)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	addedTeacher := make([]Teacher, len(NewTeachers))
	for i, newTeacher := range NewTeachers {
		newTeacher.ID = nextId
		teachers[nextId] = newTeacher
		addedTeacher[i] = newTeacher
		nextId++
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := struct {
		Status string    `json:"status"`
		Count  int       `json:"count"`
		Data   []Teacher `json:"data"`
	}{
		Status: "success",
		Count:  len(addedTeacher),
		Data:   addedTeacher,
	}
	json.NewEncoder(w).Encode(response)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Welcome to the home page") // 1 way
	w.Write([]byte("Welcome to the home page"))
}

func TeachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Calling a get function
		GetTeachersHandeler(w, r)
		return
	case http.MethodPost:
		w.Write([]byte("Welcome to Teacher POST"))
		AddTeacherHandler(w, r)
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

	// rl := mw.NewRateLimiter(5, time.Minute)
	// hpp := mw.HPPOptions{
	// 	CheckQuery:                  true,
	// 	CheckBody:                   true,
	// 	CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
	// 	Whitelist:                   []string{"sortBy", "sortOrder", "name", "age", "class"},
	// }

	// securemux := mw.CORS(rl.MiddleWare(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Compression(mw.HPP(hpp)(mux))))))
	// securemux := ApplyMiddleWares(mux, mw.HPP(hpp), mw.Compression, mw.SecurityHeaders, mw.ResponseTimeMiddleware, rl.MiddleWare, mw.CORS)
	securemux := mw.SecurityHeaders(mux)

	// Create custom server :
	server := &http.Server{
		Addr:      port,
		Handler:   securemux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port ", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatal(err)
	}
}

// Middleware is a function that wraps an http.Handler with additional functionality
type MiddleWare func(http.Handler) http.Handler

func ApplyMiddleWares(handler http.Handler, middlewares ...MiddleWare) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
