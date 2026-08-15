package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	utilssql "order_mgt/pkg/utils_sql"
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

func UploadProductImage(minioService *utilssql.MinioService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, 50<<20) // 50 MB
		err := r.ParseMultipartForm(50 << 20)
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

		objectName, err := minioService.Upload(r.Context(), file, fileheader)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res := struct {
			Success bool
			Message string
			URL     string
		}{
			Success: true,
			Message: fmt.Sprintln("File", objectName, "has been sucessfully uploaded"),
			URL:     minioService.GetURL(objectName),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&res)

	}
}
