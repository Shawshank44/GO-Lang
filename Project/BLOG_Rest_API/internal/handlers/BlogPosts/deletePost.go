package blogposts

import (
	repositories "blog_rest_api/internal/repositories/Post_SQL"
	"encoding/json"
	"net/http"
	"strconv"
)

func DeletePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idstr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Unable to conv Id", http.StatusBadRequest)
		return
	}

	err = repositories.DeletePostFromDB(r.Context(), id)
	if err != nil {
		http.Error(w, "Unable to delete the post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	res := struct {
		Success string
		ID      int
	}{
		Success: "Post Success fully deleted",
		ID:      id,
	}

	json.NewEncoder(w).Encode(&res)

}
