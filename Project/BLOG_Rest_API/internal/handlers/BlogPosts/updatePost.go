package blogposts

import (
	"blog_rest_api/internal/models"
	repositories "blog_rest_api/internal/repositories/Post_SQL"
	utilssql "blog_rest_api/internal/repositories/Utils_SQL"
	"encoding/json"
	"net/http"
	"strconv"
)

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	sessionID := r.URL.Query().Get("session_id")
	postID := r.URL.Query().Get("id")

	id, err := strconv.Atoi(postID)
	if err != nil {
		http.Error(w, "unable to convert id", http.StatusForbidden)
		return
	}

	var newPost models.Post
	err = json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	err = utilssql.FinalizeSessionUploads(r.Context(), sessionID)
	if err != nil {
		http.Error(w, "Unable to finalize uploads", http.StatusInternalServerError)
		return
	}

	err = repositories.UpdatePostInDB(r.Context(), &newPost, id)
	if err != nil {
		http.Error(w, "Unable to update the post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	res := struct {
		Success string
		Post    models.Post
	}{
		Success: "post updated sucessfully",
		Post:    newPost,
	}

	json.NewEncoder(w).Encode(&res)

}
