package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/internal/models"
	"strconv"
	"strings"
	"sync"
)

var (
	teachers = make(map[int]models.Teacher)
	mutex    = &sync.Mutex{}
	nextId   = 1
)

func init() {
	teachers[nextId] = models.Teacher{
		ID:        nextId,
		FirstName: "John",
		LastName:  "Doe",
		Class:     "9A",
		Subject:   "Math",
	}
	nextId++
	teachers[nextId] = models.Teacher{
		ID:        nextId,
		FirstName: "Jane",
		LastName:  "smith",
		Class:     "10A",
		Subject:   "Algebra",
	}
	nextId++
	teachers[nextId] = models.Teacher{
		ID:        nextId,
		FirstName: "Jane",
		LastName:  "Doe",
		Class:     "10B",
		Subject:   "Drawing",
	}
	nextId++
	teachers[nextId] = models.Teacher{
		ID:        nextId,
		FirstName: "Antheny",
		LastName:  "Missery",
		Class:     "10C",
		Subject:   "Science",
	}
	nextId++
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

func GetTeachersHandeler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idstr := strings.TrimSuffix(path, "/")
	fmt.Println(idstr)

	if idstr == "" {
		firstName := r.URL.Query().Get("first_name") // encoding/json package automatically converts your exported field names into JSON keys
		lastName := r.URL.Query().Get("last_name")

		TeacherList := make([]models.Teacher, 0, len(teachers))
		for _, teacher := range teachers {
			if (firstName == "" || teacher.FirstName == firstName) && (lastName == "" || teacher.LastName == lastName) {
				TeacherList = append(TeacherList, teacher)
			}
		}
		response := struct {
			Status string           `json:"status"`
			Count  int              `json:"count"`
			Data   []models.Teacher `json:"data"`
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
	var NewTeachers []models.Teacher
	err := json.NewDecoder(r.Body).Decode(&NewTeachers)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	addedTeacher := make([]models.Teacher, len(NewTeachers))
	for i, newTeacher := range NewTeachers {
		newTeacher.ID = nextId
		teachers[nextId] = newTeacher
		addedTeacher[i] = newTeacher
		nextId++
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Teacher `json:"data"`
	}{
		Status: "success",
		Count:  len(addedTeacher),
		Data:   addedTeacher,
	}
	json.NewEncoder(w).Encode(response)
}
