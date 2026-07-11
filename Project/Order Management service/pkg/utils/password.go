package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	memory      = 64 * 1024
	iterations  = 3
	parallelism = 4
	saltLength  = 16
	keyLength   = 32
)

func HashPassword(password string) (string, error) {
	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return "", ErrorHandler(err, "Unable to generate salt.")
	}

	hash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, keyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedhash := fmt.Sprintf("%s$%s", b64Salt, b64hash)

	return encodedhash, nil
}

func VerifyPassword(password, encodedhash string) error {
	parts := strings.Split(encodedhash, "$")
	if len(parts) != 2 {
		return ErrorHandler(errors.New("invalid encoding hash format"), "Internal server error")
	}

	b64Salt := parts[0]
	b64hash := parts[1]

	saltBytes, err := base64.RawStdEncoding.DecodeString(b64Salt)
	if err != nil {
		return ErrorHandler(err, "Unable to decode")
	}

	hashBytes, err := base64.RawStdEncoding.DecodeString(b64hash)
	if err != nil {
		return ErrorHandler(err, "Unale to decode")
	}

	hash := argon2.IDKey([]byte(password), saltBytes, iterations, memory, parallelism, keyLength)
	if len(hash) != len(hashBytes) {
		return ErrorHandler(errors.New("hash length mismatch"), "Incorrect password")
	}

	if subtle.ConstantTimeCompare(hash, hashBytes) == 1 {
		return nil
	}

	return ErrorHandler(errors.New("invalid password"), "invalid password")
}

func IsSameAsOldPassword(password, encodedHash string) (bool, error) {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 2 {
		return false, ErrorHandler(errors.New("invalid encoding hash format"), "Internal server error")
	}

	saltBytes, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return false, ErrorHandler(err, "Unable to decode")
	}

	hashBytes, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, ErrorHandler(err, "Unable to decode")
	}

	hash := argon2.IDKey([]byte(password), saltBytes, iterations, memory, parallelism, keyLength)

	if len(hash) != len(hashBytes) {
		return false, ErrorHandler(errors.New("hash length mismatch"), "hash length mismatch")
	}

	return subtle.ConstantTimeCompare(hash, hashBytes) == 1, nil
}
