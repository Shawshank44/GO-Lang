package auth

import "fmt"

// if you want to export the function just make first letter capital
func LoginWithCredentials(username string, password string) {
	fmt.Println("Logging user", username)
}
