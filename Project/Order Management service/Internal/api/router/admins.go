package router

import (
	"fmt"
	"net/http"
	"order_mgt/Internal/api/handlers"
)

func AdminRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Read :
	mux.HandleFunc("GET /api/admin/super/getadmins", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to getadmin page"))
	})
	mux.HandleFunc("GET /api/admin/super/getadmin/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to get admins by id page : ", r.PathValue("id"))
	})

	// Create :
	mux.HandleFunc("POST /api/admin/super/register", handlers.RegisterAdmin)
	mux.HandleFunc("POST /api/admin/super/login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to admin login page"))
	})
	mux.HandleFunc("POST /api/admin/super/logout", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to admin logout page"))
	})

	// Update :
	mux.HandleFunc("PATCH /api/admin/super/updateadmin/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to admin update page"))
	})

	// Delete :
	mux.HandleFunc("DELETE /api/admin/super/deactivate", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to admin deactivate page"))
	})
	mux.HandleFunc("DELETE /api/admin/super/deactivateusers", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to admin deactivateusers page"))
	})

	// MISCS :
	mux.HandleFunc("POST /api/admin/super/forgotpassword", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to admin forgotpassword page"))
	})
	mux.HandleFunc("POST /api/admin/super/resetpassword", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to admin resetpassword page"))
	})

	return mux
}
