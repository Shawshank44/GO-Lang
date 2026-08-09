package main

import (
	"log"
	"net/http"
	"order_mgt/Internal/api/middlewares"
	"order_mgt/Internal/api/router"
	sqlconnect "order_mgt/Internal/repository/sqlConnect"
	utilssql "order_mgt/pkg/utils_sql"
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

	Min, err := utilssql.NewMinioService()
	if err != nil {
		log.Fatal("failed to start the Minio serve", err)
	}

	err = Min.CheckBucket()
	if err != nil {
		log.Fatal("Invalid Bucket", err)
	}

	routers := router.MainRouter(Min)
	jwtMiddlewares := middlewares.MiddlewaresExcludeParts(middlewares.JWTMiddleware, "/api/admin/super/register", "/api/admin/super/login", "/api/admin/super/forgotpassword", "/api/admin/super/resetpassword")
	securemux := middlewares.ApplyMiddleWares(routers, jwtMiddlewares)

	server := &http.Server{
		Addr:    os.Getenv("API_PORT"),
		Handler: securemux,
	}

	log.Println("MINIO Service successfully connected..")

	log.Printf("Server started successfully on http://localhost%s", os.Getenv("API_PORT"))
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalln("Server error : ", err)
	}
}
