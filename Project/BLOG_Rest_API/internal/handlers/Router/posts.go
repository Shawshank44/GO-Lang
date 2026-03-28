package router

import (
	blogposts "blog_rest_api/internal/handlers/BlogPosts"
	"net/http"
)

func PostsRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /session/create", blogposts.CreateSession)

	mux.HandleFunc("POST /assets/uploads/", blogposts.Uploader)
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir(`C:\Users\Shashank.BR\OneDrive\Desktop\Go programing\Project\BLOG_Rest_API\cmd\server\uploads\`))))

	mux.HandleFunc("GET /getposts", blogposts.GetPosts)
	mux.HandleFunc("GET /myposts", blogposts.MyPosts)
	mux.HandleFunc("GET /getpost/{id}", blogposts.GetPost)

	mux.HandleFunc("POST /createpost", blogposts.CreatePost)

	mux.HandleFunc("PATCH /updatepost", blogposts.UpdatePost)

	mux.HandleFunc("DELETE /deletepost", blogposts.DeletePost)

	return mux
}
