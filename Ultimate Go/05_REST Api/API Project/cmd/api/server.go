package main

import (
	"crypto/tls"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	mw "restapi/internal/api/middlewares"
	"restapi/internal/api/router"
	"restapi/pkg/utils"
	"time"

	"github.com/joho/godotenv"
)

//go:embed .env
var envFile embed.FS

func loadEnvFromEmbededFile() {
	// Read the embeded .env file
	content, err := envFile.ReadFile(".env")
	if err != nil {
		log.Fatalf("Error reading the .env file : %v", err)
		return
	}

	// Creating a temp file to load the env variables
	tempfile, err := os.CreateTemp("", ".env")
	if err != nil {
		log.Fatalf("Error creating the temporary .env file : %v", err)
		return
	}
	defer os.Remove(tempfile.Name())

	// Write the env variables in the temp .env file
	_, err = tempfile.Write(content)
	if err != nil {
		log.Fatalf("Error writing the temporary .env file : %v", err)
		return
	}

	err = tempfile.Close()
	if err != nil {
		log.Fatalf("Error closing the temporary .env file : %v", err)
		return
	}

	// Load env vars from the temp file
	err = godotenv.Load(tempfile.Name())
	if err != nil {
		log.Fatalf("Error loading the temporary .env file : %v", err)
		return
	}
}

func main() {
	//  Only in Development, for running the source code
	// err := godotenv.Load()
	// if err != nil {
	// 	return
	// }

	// load environment variables from the embedded .env
	loadEnvFromEmbededFile()
	fmt.Println("ENvironment variable CERT_FILE : ", os.Getenv("CERT_FILE"))

	port := os.Getenv("API_PORT")

	// cert := "cert.pem"
	// key := "key.pem"

	cert := os.Getenv("CERT_FILE")
	key := os.Getenv("KEY_FILE")

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS10,
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
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatal(err)
	}
}
