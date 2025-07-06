package main

import (
	"log"

	"go.uber.org/zap"
)

// Creating a customer logger
// var (
// 	infoLogger = log.New(os.Stdout, "INFO : ", log.Ldate|log.Ltime|log.Lshortfile)
// 	warnLogger = log.New(os.Stdout, "WARN : ", log.Ldate|log.Ltime|log.Lshortfile)
// 	errLogger  = log.New(os.Stdout, "ERROR : ", log.Ldate|log.Ltime|log.Lshortfile)
// )

func main() {
	//printing a log message
	// log.Println("This is log message")

	// // Setting a prefix
	// log.SetPrefix("INFO : ")
	// log.Println("This is a info message")

	// // Log flags :
	// log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// log.Println("This is a log message with only date time and file name")

	// // Using the customer loggers :
	// infoLogger.Println("This is Info Message")
	// warnLogger.Println("This is Warn Message")
	// errLogger.Println("This is Error Message")

	// // Logging errors in a separate file :
	// file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalf("Failed to open log file : %v", err)
	// }
	// defer file.Close()

	// debugLogger := log.New(file, "debug : ", log.Ldate|log.Ltime|log.Lshortfile)
	// debugLogger1 := log.New(file, "debug-WARN : ", log.Ldate|log.Ltime|log.Lshortfile)
	// debugLogger2 := log.New(file, "debug-ERROR : ", log.Ldate|log.Ltime|log.Lshortfile)

	// debugLogger.Println("This is Debug message")
	// debugLogger1.Println("This is a WARNING ")
	// debugLogger2.Println("This is a ERROR")

	// Using Logrus 3rd party package :
	// logs := logrus.New()

	// logs.SetLevel(logrus.InfoLevel)

	// logs.SetFormatter(&logrus.JSONFormatter{})

	// // Example :
	// logs.Info("This is an Info Message. ")
	// logs.Warn("This is a Warning Message.")
	// logs.Error("This is a error Message. ")

	// logs.WithFields(logrus.Fields{
	// 	"username": "John doe",
	// 	"Method":   "GET",
	// }).Info("User logged in")

	// Using Zap 3rd part package :
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("Error occured")
	}

	defer logger.Sync()

	logger.Info("This is an info message")
	logger.Info("User logged in ", zap.String("Username", "John doe"), zap.String("Method", "GET"))

}
