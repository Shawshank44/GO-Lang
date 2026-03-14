package blogposts

import (
	"blog_rest_api/internal/middlewares"
	"blog_rest_api/internal/models"
	repositories "blog_rest_api/internal/repositories/Post_SQL"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	posts, err := repositories.GetPostsFromDB(r.Context())
	if err != nil {
		http.Error(w, "Unable to fetch the post", http.StatusInternalServerError)
		return
	}

	res := struct {
		Status string
		Posts  []models.Post
	}{
		Status: "Success",
		Posts:  *posts,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func MyPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username, ok := r.Context().Value(middlewares.UsernameKey).(string)
	if !ok {
		http.Error(w, "username not found in context", http.StatusUnauthorized)
		return
	}

	posts, err := repositories.MyPostsFromDB(r.Context(), username)
	if err != nil {
		http.Error(w, "Unable to fetch posts ", http.StatusInternalServerError)
		return
	}

	res := struct {
		Status string
		Posts  []models.Post
	}{
		Status: "Success",
		Posts:  *posts,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

func GetPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idstr := r.PathValue("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Inavlid id type", http.StatusBadRequest)
		return
	}

	post, err := repositories.GetPostFromDB(r.Context(), id)

	if err != nil {
		http.Error(w, "Unable to fetch the post", http.StatusInternalServerError)
		return
	}

	res := struct {
		Status string
		Posts  models.Post
	}{
		Status: "Success",
		Posts:  *post,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}
