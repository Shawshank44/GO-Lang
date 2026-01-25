package main

import (
	router "blog_rest_api/internal/handlers/Router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	routers := router.MainRouter()

	fmt.Println("Server successfully created on http://localhost:8080")
	err := http.ListenAndServe(":8080", routers)
	if err != nil {
		log.Fatal(err.Error())
	}
}
