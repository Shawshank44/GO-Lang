package utils

import (
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


func SignToken(userID string, Username, Role string) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	jwtExpiresIn := os.Getenv("JWT_EXPIRES_IN")
	claims := jwt.MapClaims{
		"uid":  userID,
		"user": Username,
		"role": Role,
	}

	if jwtExpiresIn != "" {
		duration, err := time.ParseDuration(jwtExpiresIn)
		if err != nil {
			return "", ErrorHandler(err, "internal error")
		}
		claims["exp"] = jwt.NewNumericDate(time.Now().Add(duration))
	} else {
		claims["exp"] = jwt.NewNumericDate(time.Now().Add(15 * time.Minute))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte((jwtSecret)))
	if err != nil {
		return "", ErrorHandler(err, "Internal error")
	}

	return signedToken, nil
}

var JwtStore = JWTStore{
	Tokens: make(map[string]time.Time),
}

type JWTStore struct {
	Mu     sync.Mutex
	Tokens map[string]time.Time
}

func (store *JWTStore) Addtoken(token string, expiryTime time.Time) {
	store.Mu.Lock()
	defer store.Mu.Unlock()

	store.Tokens[token] = expiryTime
}

func (store *JWTStore) CleanUpExpiredToken() {
	for {
		time.Sleep(2 * time.Minute)

		store.Mu.Lock()
		for token, timestamp := range store.Tokens {
			if time.Now().After(timestamp) {
				delete(store.Tokens, token)
			}
		}
		store.Mu.Unlock()
	}
}

func (store *JWTStore) IsLoggedOut(token string) bool {
	store.Mu.Lock()
	defer store.Mu.Unlock()
	_, ok := store.Tokens[token]
	return ok
}
