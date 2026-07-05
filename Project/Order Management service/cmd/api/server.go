package main

import (
	"log"
	"net/http"
	"order_mgt/Internal/api/middlewares"
	"order_mgt/Internal/api/router"
	sqlconnect "order_mgt/Internal/repository/sqlConnect"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	_, err = sqlconnect.ConnectDB()
	if err != nil {
		log.Fatalln("unable to connect to DB", err)
	}

	routers := router.MainRouter()

	securemux := middlewares.ApplyMiddleWares(routers)

	server := &http.Server{
		Addr:    os.Getenv("API_PORT"),
		Handler: securemux,
	}

	log.Printf("Server started successfully on http://localhost%s", os.Getenv("API_PORT"))
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalln("Server error : ", err)
	}
}
