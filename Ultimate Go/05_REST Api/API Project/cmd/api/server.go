package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	mw "restapi/internal/api/middlewares"
	"restapi/internal/api/router"
	"restapi/internal/repository/sqlconnect"
	"restapi/pkg/utils"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		return
	}

	_, err = sqlconnect.ConnectDB()
	if err != nil {
		utils.ErrorHandler(err, "")
		return
	}

	port := os.Getenv("API_PORT")

	cert := "cert.pem"
	key := "key.pem"

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	rl := mw.NewRateLimiter(5, time.Minute)
	hpp := mw.HPPOptions{
		CheckQuery:                  true,
		CheckBody:                   true,
		CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
		Whitelist:                   []string{"sortBy", "sortOrder", "name", "age", "class"},
	}

	// securemux := mw.CORS(rl.MiddleWare(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Compression(mw.HPP(hpp)(mux))))))
	// securemux := jwtMiddlewares(mw.SecurityHeaders(router.MainRouter()))
	// securemux := mw.SecurityHeaders(router.MainRouter())
	// securemux := mw.XSSMiddleWares(router.MainRouter())
	router := router.MainRouter()
	jwtMiddlewares := mw.MiddlewaresExcludeParts(mw.JWTMiddlewares, "/execs/login", "/execs/forgotpassword", "/execs/resetpassword/reset")
	securemux := utils.ApplyMiddleWares(router, mw.SecurityHeaders, mw.Compression, mw.HPP(hpp), mw.XSSMiddleWares, jwtMiddlewares, mw.ResponseTimeMiddleware, rl.MiddleWare, mw.CORS)

	// Create custom server :
	server := &http.Server{
		Addr:      port,
		Handler:   securemux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port ", port)
	err = server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatal(err)
	}
}
