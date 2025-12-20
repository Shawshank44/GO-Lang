package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func HashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hashs := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hashs[:])
}

func main() {
	password := "password1234"

	// // SHA-256 :
	// hash256 := sha256.Sum256([]byte(password))
	// fmt.Println(hash256) // returns a byte slice of a hash
	// fmt.Printf("SHA-256 hash hex value : %x \n", hash256)

	// // SHA-512 :
	// hash512 := sha512.Sum512([]byte(password))
	// fmt.Println(hash512) // returns a byte slice of a hash
	// fmt.Printf("SHA-512 has hex value : %x \n", hash512)

	// Salting :
	salt, err := GenerateSalt()
	fmt.Printf("Actual salt value %x \n", salt)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Hashing password
	hashed := HashPassword(password, salt)

	// printing the salt and password
	saltstr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt : ", saltstr)
	fmt.Println("Hashed password : ", hashed)

	// Verifying the password :
	deCodedSalt, err := base64.StdEncoding.DecodeString(saltstr)
	if err != nil {
		fmt.Println(err)
	}
	checkhash := HashPassword(password, deCodedSalt)

	// Compare the stored hash
	if hashed == checkhash {
		fmt.Println("Password is correct")
	} else {
		fmt.Println("Invalid Submission")
	}
}
