package main

import (
	"blog_rest_api/internal/config"
	"blog_rest_api/internal/db"
	router "blog_rest_api/internal/handlers/Router"
	"blog_rest_api/internal/middlewares"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	routers := router.MainRouter()
	jwtMiddlewares := middlewares.MiddleWaresExcludeRoutes(middlewares.JWTMiddleware, "/users/register", "/users/login", "/users/forgotpassword", "/users/resetpassword")
	securemux := middlewares.ApplyMiddleWares(routers, jwtMiddlewares)

	server := &http.Server{
		Addr:    cfg.API_PORT,
		Handler: securemux,
	}

	fmt.Printf("Server successfully created on http://localhost%s", cfg.API_PORT)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
