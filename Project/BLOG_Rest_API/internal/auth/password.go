package auth

import (
	"blog_rest_api/pkg/utils"
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
		return "", utils.ErrorHandler(err, "Unable to generate salt.")
	}

	hash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, keyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("%s$%s", b64Salt, b64hash)
	return encodedHash, nil
}

func VerifyPassword(password, encodedHash string) error {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 2 {
		return utils.ErrorHandler(errors.New("invalid encoded hash format"), "Internal server error")
	}
	b64Salt := parts[0]
	b64hash := parts[1]

	saltBytes, err := base64.RawStdEncoding.DecodeString(b64Salt)
	if err != nil {
		return utils.ErrorHandler(err, "Unable to Decode")
	}

	hashBytes, err := base64.RawStdEncoding.DecodeString(b64hash)
	if err != nil {
		return utils.ErrorHandler(err, "Unable to decode")
	}

	hash := argon2.IDKey([]byte(password), saltBytes, iterations, memory, parallelism, keyLength)

	if len(hash) != len(hashBytes) {
		return utils.ErrorHandler(errors.New("hash length mismatch"), "Incorrect password")
	}

	if subtle.ConstantTimeCompare(hash, hashBytes) == 1 {
		return nil
	}

	return utils.ErrorHandler(errors.New("invalid password"), "invalid password")
}
