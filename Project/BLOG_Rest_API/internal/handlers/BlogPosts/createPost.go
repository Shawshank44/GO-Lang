package blogposts

import (
	"blog_rest_api/internal/middlewares"
	"blog_rest_api/internal/models"
	repositories "blog_rest_api/internal/repositories/Post_SQL"
	utilssql "blog_rest_api/internal/repositories/Utils_SQL"
	"encoding/json"
	"net/http"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sessionID := r.URL.Query().Get("session_id")

	var post models.Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	content, err := json.Marshal(post.Content)
	if err != nil {
		http.Error(w, "Invalid content body", http.StatusBadRequest)
		return
	}

	tags, err := json.Marshal(post.Tags)
	if err != nil {
		http.Error(w, "Invalid Tags", http.StatusBadRequest)
		return
	}

	username, ok := r.Context().Value(middlewares.UsernameKey).(string)
	if !ok {
		http.Error(w, "username not found in context", http.StatusUnauthorized)
		return
	}

	err = repositories.CreatePostInDB(r.Context(), username, post.Title, content, tags)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utilssql.FinalizeSessionUploads(r.Context(), sessionID)

	w.Header().Set("Content-Type", "application/json")
	res := struct {
		Success string
		Post    models.Post
	}{
		Success: "Post has been created",
		Post:    post,
	}

	json.NewEncoder(w).Encode(&res)

}
