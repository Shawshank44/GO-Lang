package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	utilssql "order_mgt/pkg/utils_sql"

	"github.com/google/uuid"
)

func isAllowedType(contentType string) bool {
	allowed := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/webp": true,
		"video/mp4":  true,
	}
	return allowed[contentType]
}

func CreateSession(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sessionID := uuid.New().String()

	err := utilssql.CreateSessionsInDB(r.Context(), sessionID)
	if err != nil {
		http.Error(w, "Internal server error unable to create session", http.StatusInternalServerError)
		return
	}

	res := struct {
		Success   bool   `json:"success"`
		SessionID string `json:"session_id"`
	}{
		Success:   true,
		SessionID: sessionID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&res)
}

func UploadProductImage(minioService *utilssql.MinioService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		sessionID := r.URL.Query().Get("session_id")

		exists, err := utilssql.ValidateSessionsInDB(r.Context(), sessionID)
		if err != nil {
			http.Error(w, "Unable to find the session", http.StatusInternalServerError)
			return
		}

		if !exists {
			http.Error(w, "Invalid session", http.StatusForbidden)
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, 50<<20) // 50 MB
		err = r.ParseMultipartForm(50 << 20)
		if err != nil {
			http.Error(w, "File is too large for uploaded", http.StatusBadRequest)
			return
		}

		file, fileheader, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "file is required", http.StatusBadRequest)
			return
		}

		defer file.Close()

		buffer := make([]byte, 512)
		_, err = file.Read(buffer)
		if err != nil {
			http.Error(w, "Invalid buffer space", http.StatusInternalServerError)
			return
		}

		contentType := http.DetectContentType(buffer)
		if !isAllowedType(contentType) {
			http.Error(w, "Unsupported file type", http.StatusBadRequest)
			return
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			http.Error(w, "unable to seek", http.StatusBadRequest)
			return
		}

		ObjectName, err := minioService.Upload(r.Context(), file, fileheader)
		if err != nil {
			http.Error(w, "Unable to upload file", http.StatusInternalServerError)
			return
		}

		err = utilssql.UploadToDB(r.Context(), sessionID, ObjectName)
		if err != nil {
			delerr := minioService.Delete(r.Context(), ObjectName)
			if delerr != nil {
				log.Printf("failed to rollback MinIO object %q: %v\n", ObjectName, delerr)
			}
			http.Error(w, "Unable to save uploaded file", http.StatusInternalServerError)
			return
		}

		res := struct {
			Success bool
			Message string
			URL     string
		}{
			Success: true,
			Message: fmt.Sprintf("File %s has been successfully uploaded", ObjectName),
			URL:     minioService.GetURL(ObjectName),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&res)

	}
}
