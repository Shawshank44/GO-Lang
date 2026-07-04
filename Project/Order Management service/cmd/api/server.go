package main

import (
	"fmt"
	"log"
	"net/http"
	"order_mgt/Internal/api/middlewares"
	"order_mgt/Internal/api/router"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	routers := router.MainRouter()

	securemux := middlewares.ApplyMiddleWares(routers)

	server := &http.Server{
		Addr:    os.Getenv("API_PORT"),
		Handler: securemux,
	}

	fmt.Printf("Server started successfully on http://localhost%s", os.Getenv("API_PORT"))
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalln("Server error : ", err)
	}
}
