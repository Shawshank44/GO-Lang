package utils

import (
	"fmt"
	"log"
	"os"
)

var errLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

func ErrorHandler(err error, message string) error {
	if err == nil {
		return nil
	}
	errLogger.Printf("%s: %v", message, err)

	return fmt.Errorf("%s: %w", message, err)
}
