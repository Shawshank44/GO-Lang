package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/internal/models"
	"restapi/internal/repository/sqlconnect"
	"strconv"
	"strings"
	"sync"
)

var (
	teachers = make(map[int]models.Teacher)
	mutex    = &sync.Mutex{}
)

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
	db, err := sqlconnect.ConnectDB()

	if err != nil {
		http.Error(w, "Error in connecting to Database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idstr := strings.TrimSuffix(path, "/")
	fmt.Println(idstr)

	if idstr == "" {
		firstName := r.URL.Query().Get("first_name") // encoding/json package automatically converts your exported field names into JSON keys
		lastName := r.URL.Query().Get("last_name")

		query := "SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE 1=1"

		var args []interface{}

		if firstName != "" {
			query += " AND first_name = ?"
			args = append(args, firstName)
		}
		if lastName != "" {
			query += " AND last_name = ?"
			args = append(args, lastName)
		}

		rows, err := db.Query(query, args...)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Database Query error", http.StatusInternalServerError)
			return
		}

		defer rows.Close()

		TeacherList := make([]models.Teacher, 0)
		for rows.Next() {
			var teacher models.Teacher
			err = rows.Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
			if err != nil {
				http.Error(w, "Error scanning database results", http.StatusInternalServerError)
			}
			TeacherList = append(TeacherList, teacher)
		}
		response := struct {
			Status string           `json:"status"`
			Count  int              `json:"count"`
			Data   []models.Teacher `json:"data"`
		}{
			Status: "Success",
			Count:  len(TeacherList),
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
	var teacher models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", id).Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
	if err == sql.ErrNoRows {
		http.Error(w, "Teacher not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "database query error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teacher)
}

func AddTeacherHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sqlconnect.ConnectDB()

	if err != nil {
		http.Error(w, "Error in connecting to Database", http.StatusInternalServerError)
		return
	}

	defer db.Close()

	var NewTeachers []models.Teacher
	err = json.NewDecoder(r.Body).Decode(&NewTeachers)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("INSERT INTO teachers (first_name, last_name, email, class, subject) VALUES (?,?,?,?,?)")
	if err != nil {
		http.Error(w, "Error in Preparing SQL query", http.StatusInternalServerError)
	}

	defer stmt.Close()

	addedTeachers := make([]models.Teacher, len(NewTeachers))
	for i, Newteacher := range NewTeachers {
		res, err := stmt.Exec(Newteacher.FirstName, Newteacher.LastName, Newteacher.Email, Newteacher.Class, Newteacher.Subject)
		if err != nil {
			http.Error(w, "error inserting data in database", http.StatusInternalServerError)
			return
		}
		lastID, err := res.LastInsertId()
		if err != nil {
			http.Error(w, "Error getting last insert ID", http.StatusInternalServerError)
		}
		Newteacher.ID = int(lastID)
		addedTeachers[i] = Newteacher
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Teacher `json:"data"`
	}{
		Status: "success",
		Count:  len(addedTeachers),
		Data:   addedTeachers,
	}
	json.NewEncoder(w).Encode(response)
}
