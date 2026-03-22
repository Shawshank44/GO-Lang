package blogposts

import (
	repositories "blog_rest_api/internal/repositories/Post_SQL"
	utilssql "blog_rest_api/internal/repositories/Utils_SQL"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

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
	sessionID := uuid.New().String()

	err := utilssql.CreateSessionInDB(r.Context(), sessionID)
	if err != nil {
		http.Error(w, "Internal server error unable to create session", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"session_id": sessionID,
	})
}

func Uploader(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sessionID := r.URL.Query().Get("session_id")

	var fileURL string

	exists, err := utilssql.ValidateSession(r.Context(), sessionID)
	if err != nil {
		http.Error(w, "Unable to find the session", http.StatusInternalServerError)
		return
	}

	if exists {
		r.Body = http.MaxBytesReader(w, r.Body, 50<<20) // 50 MB

		err = r.ParseMultipartForm(50 << 20)
		if err != nil {
			http.Error(w, "File is too large for upload", http.StatusBadRequest)
			return
		}

		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Invalid file", http.StatusBadRequest)
			return
		}

		defer file.Close()

		buffer := make([]byte, 512)
		_, err = file.Read(buffer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
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

		fileName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), header.Filename)

		uploadPath := `C:\Users\Shashank.BR\OneDrive\Desktop\Go programing\Project\BLOG_Rest_API\cmd\server\uploads\` + fileName

		dst, err := os.Create(uploadPath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(w, "Save failed ", http.StatusInternalServerError)
			return
		}

		fileURL = "/uploads/" + fileName

		err = repositories.UploadToDB(r.Context(), sessionID, fileURL)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Invalid session", http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "**UPLOADED**",
		"URL":    fileURL,
	})
}
