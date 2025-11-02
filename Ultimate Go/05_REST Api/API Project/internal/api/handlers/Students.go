package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"restapi/internal/models"
	"restapi/internal/repository/sqlconnect"
	"restapi/pkg/utils"
	"strconv"
)

// Get Students
func GetStudentsHandeler(w http.ResponseWriter, r *http.Request) {
	var students []models.Student
	page, limit := utils.GetPaginationParams(r)

	students, totalStudents, err := sqlconnect.GetStudentsDBHandeler(students, r, limit, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	totalPages := (totalStudents + limit - 1) / limit
	response := struct {
		Status     string           `json:"status"`
		Count      int              `json:"count"`
		TotalPages int              `json:"total_pages"`
		PageNo     int              `json:"page_no"`
		PageSize   int              `json:"page_limit"`
		Data       []models.Student `json:"data"`
	}{
		Status:     "Success",
		Count:      totalStudents,
		TotalPages: totalPages,
		PageNo:     page,
		PageSize:   limit,
		Data:       students,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

// Get Student By ID
func GetStudentHandeler(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	student, err := sqlconnect.GetStudentDBHandeler(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

// POST Students
func AddStudentHandler(w http.ResponseWriter, r *http.Request) {
	var NewStudents []models.Student
	var RawStudents []map[string]interface{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	err = json.Unmarshal(body, &RawStudents)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	fields := utils.GetFieldNames(models.Student{})

	allowedFields := make(map[string]struct{})
	for _, field := range fields {
		allowedFields[field] = struct{}{}
	}

	for _, student := range RawStudents {
		for key := range student {
			_, ok := allowedFields[key]
			if !ok {
				http.Error(w, "Unacceptable fields found in request. Only use allowed fields", http.StatusBadRequest)
				return
			}
		}
	}

	err = json.Unmarshal(body, &NewStudents)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	for _, student := range NewStudents {
		err = utils.CheckBlankFields(student)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	addedStudents, err := sqlconnect.POSTStudentDBHandler(NewStudents)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Student `json:"data"`
	}{
		Status: "success",
		Count:  len(addedStudents),
		Data:   addedStudents,
	}
	json.NewEncoder(w).Encode(response)
}

// PUT Method By ID (expects complete replacement of the data)
func PutStudentHandler(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid student request", http.StatusBadRequest)
		return
	}

	var UpdatedStudent models.Student
	err = json.NewDecoder(r.Body).Decode(&UpdatedStudent)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	UpdatedStudentFromDB, err := sqlconnect.PutStudentDBHandler(id, UpdatedStudent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(UpdatedStudentFromDB)
}

// PATCH Method (expects replacement of required data)
func PatchStudentsHandler(w http.ResponseWriter, r *http.Request) {
	var updates []map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = sqlconnect.PatchStudentsDBHandler(updates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// PATCH Method by ID (expects replacement of required data)
func PatchStudentHandler(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid student request", http.StatusBadRequest)
		return
	}
	var Updates map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&Updates)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}
	UpdatedStudent, err := sqlconnect.PatchStudentDBHandler(id, Updates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(UpdatedStudent)

}

// DELETE Method
func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid student request", http.StatusBadRequest)
		return
	}
	err = sqlconnect.DeleteStudentDBHandler(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// w.WriteHeader(http.StatusNoContent)
	// Response Body
	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Status string `json:"status"`
		ID     int    `json:"id"`
	}{
		Status: "Student Sucessfully deleted",
		ID:     id,
	}

	json.NewEncoder(w).Encode(response)

}

func DeleteStudentsHandler(w http.ResponseWriter, r *http.Request) {
	var ids []int
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	deletedIds, err := sqlconnect.DeleteStudentsDBHandler(ids)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Status    string `json:"status"`
		DeletedID []int  `json:"deleted_id"`
	}{
		Status:    "Techer Successfully Deleted",
		DeletedID: deletedIds,
	}

	json.NewEncoder(w).Encode(response)
}
